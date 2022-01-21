import Button from "@mui/material/Button";
import Card from "@mui/material/Card";
import CardActions from "@mui/material/CardActions";
import CardContent from "@mui/material/CardContent";
import Typography from "@mui/material/Typography";
import React from "react";
import { Link } from "react-router-dom";

function NoteCard(props) {
  const { id, title, detail } = props.note;
  return (
    <Card sx={{ minWidth: 275, m: 1 }}>
      <CardContent>
        <Typography sx={{ fontSize: 14 }} color="text.secondary" gutterBottom>
          {title}
        </Typography>
        <Typography variant="h5" component="div">
          {detail}
        </Typography>
      </CardContent>
      <CardActions>  
          <Link to={`${id}`} style={{ textDecoration: "none" }}>
            <Button variant="outlined">Learn More</Button>
          </Link>
      </CardActions>
    </Card>
  );
}

export default NoteCard;