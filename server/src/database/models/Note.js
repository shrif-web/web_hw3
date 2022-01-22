import { DataTypes, Model } from "sequelize";
import sequelize from "../../utils/sequelize";
import cacheClient from "../../utils/cache-client";

class Note extends Model {
  static CACHE_NAMESPACE = "notes:";

  static async findByPk(noteId) {
    const key = `${this.CACHE_NAMESPACE}${noteId}`;
    try {
      const cachedNote = await cacheClient.getKey().sendMessage({ key });
      return new Note(JSON.parse(cachedNote.value));
    } catch (errors) {
      console.log("got from database");

      const note = await super.findByPk(noteId);
      if (note) {
        await cacheClient.setKey().sendMessage({ key, value: JSON.stringify(note) });
      }
      return note;
    }
  }

  static async create(data) {
    const newNote = await super.create(data);
    console.log("create in database");

    const key = `${this.CACHE_NAMESPACE}${newNote.id}`;
    try {
      await cacheClient.setKey().sendMessage({ key, value: JSON.stringify(newNote) });
      console.log("create in cache");
    } catch (error) {
      console.log("cannot create in cache");
    }

    return newNote;
  }

  static async update(updatedValues, query) {
    const [isUpdate, updatedNotes] = await super.update(updatedValues, query);
    if (isUpdate) {
      console.log("update in database");
      const key = `${this.CACHE_NAMESPACE}${query.where.id}`;
      try {
        await cacheClient.setKey().sendMessage({ key, value: JSON.stringify(updatedNotes[0]) });
        console.log("update in cache");
      } catch (error) {
        console.log("cannot update in cache")
      }
    }
    return [isUpdate, updatedNotes];
  }

  static async destroy(query) {
    const isDeleted = await super.destroy(query);
    if (isDeleted) {
      const key = `${this.CACHE_NAMESPACE}${query.where.id}`;
      console.log("delete from database");
      try {
        await cacheClient.remove().sendMessage({ key });
        console.log("delete from cache");
      } catch (error) {
        console.log("cannot remove from cache")
      }
    }
    return isDeleted;
  }
}

Note.init(
  {
    title: {
      type: DataTypes.STRING(256),
      allowNull: false,
    },
    detail: {
      type: DataTypes.STRING(512),
      allowNull: false,
    },
  },
  { sequelize }
);

export default Note;
