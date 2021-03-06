import Router from "@koa/router";
import validateNoteIdParams from "./validators/validate-note-id-params";
import { User, Note } from "./database/models";
import validateBody from "./validators/validate-body";
import { NOTE_POST_REQUEST_SCHEMA, NOTE_PUT_REQUEST_SCHEMA, USER_REGISTER_AND_LOGIN_REQUEST_BODY } from "./validators/schemas";
import createJwtToken from "./utils/create-jwt-token";
import isAuthenticated from "./validators/is-authenticated";
import hasPermission from "./validators/has-permission";
import RESPONSE_CODE from "./utils/response-codes";

const router = new Router();

router.post("/user/login", validateBody(USER_REGISTER_AND_LOGIN_REQUEST_BODY), async (ctx) => {
  const {
    request: {
      body: { username, password },
    },
  } = ctx;

  const user = await User.findOne({ where: { username } });

  if (await user?.isValidPassword(password)) {
    ctx.body = { message: "you have logged in successfully", token: createJwtToken(user) };
  } else {
    ctx.unauthorized({ errors: ["username and password do not match!"] });
  }
});

router.post("/user/register", validateBody(USER_REGISTER_AND_LOGIN_REQUEST_BODY), async (ctx) => {
  const {
    request: {
      body: { username, password },
    },
  } = ctx;

  const user = await User.findOne({ where: { username } });
  if (user) {
    ctx.send(RESPONSE_CODE.CONFLICT, { errors: ["username has already been taken!"] });
  } else {
    const newUser = await User.create({ username, password, isSuperuser: false });
    ctx.body = { message: "user has been created successfully", token: createJwtToken(newUser) };
  }
});

router.get("/notes", isAuthenticated(), async (ctx) => {
  const { user } = ctx;
  const whereClause = user.isSuperuser ? {} : { where: { authorId: user.userId } };

  ctx.body = await Note.findAll({...whereClause, attributes: ["id", "title", "detail"] });
});

router.get("/notes/:noteId", isAuthenticated(), validateNoteIdParams(), async (ctx) => {
  const {
    params: { noteId },
    user,
  } = ctx;
  const note = await Note.findByPk(noteId);
  if (note) {
    if (hasPermission(user, note)) {
      ctx.body = note;
    } else {
      ctx.forbidden({ errors: ["you have no permission to access data!"] });
    }
  } else {
    ctx.notFound({ errors: [`note with id ${noteId} not found!`] });
  }
});

router.post("/notes/new", isAuthenticated(), validateBody(NOTE_POST_REQUEST_SCHEMA), async (ctx) => {
  const {
    user: { userId },
    request: {
      body: { title, detail },
    },
  } = ctx;
  const newNote = await Note.create({ title, detail, authorId: userId });
  ctx.created(newNote);
});

router.put("/notes/:noteId", isAuthenticated(), validateNoteIdParams(), validateBody(NOTE_PUT_REQUEST_SCHEMA), async (ctx) => {
  const {
    params: { noteId },
    user,
    request: {
      body: { title, detail },
    },
  } = ctx;
  const updatedValues = {};
  if (title) {
    updatedValues.title = title;
  }
  if (detail) {
    updatedValues.detail = detail;
  }
  const whereClause = { id: noteId };
  if (!user.isSuperuser) {
    whereClause.authorId = user.userId;
  }

  const [isUpdate, updatedNotes] = await Note.update(updatedValues, { where: { ...whereClause }, returning: true });
  if (isUpdate) {
    [ctx.body] = updatedNotes;
  } else {
    ctx.notFound({ errors: [`note with id ${noteId} not found!`] });
  }
});

router.delete("/notes/:noteId", isAuthenticated(), validateNoteIdParams(), async (ctx) => {
  const {
    params: { noteId },
    user,
  } = ctx;

  const whereClause = { id: noteId };
  if (!user.isSuperuser) {
    whereClause.authorId = user.userId;
  }
  const isDeleted = await Note.destroy({ where: { ...whereClause } });
  if (isDeleted) {
    ctx.noContent();
  } else {
    ctx.notFound({ errors: [`note with id ${noteId} not found!`] });
  }
});

export default router;
