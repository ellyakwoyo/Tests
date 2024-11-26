import React from "react";
import {
  List,
  ListItem,
  ListItemText,
  ListItemAvatar,
  Avatar,
  Button,
  Paper,
} from "@mui/material";

const BookList = ({ books, addBookToReadingList, showAlert }) => {
  return (
    <Paper
      elevation={3}
      style={{
        maxHeight: "400px",
        overflowY: "auto",
        padding: "10px",
        marginBottom: "20px",
      }}
    >
      <List style={{ width: "100%" }}>
        {books.map((book, index) => (
          <ListItem key={index} divider>
            <ListItemAvatar>
              <Avatar alt={book.title} src={`${book.coverPhotoURL}`} />
            </ListItemAvatar>
            <ListItemText primary={book.title} secondary={book.author} />
            <Button
              variant="contained"
              color="primary"
              onClick={() => {
                addBookToReadingList(book);
                showAlert(`Added ${book.title} to the reading list`);
              }}
            >
              Add to Reading List
            </Button>
          </ListItem>
        ))}
      </List>
    </Paper>
  );
};

export default BookList;
