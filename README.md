# phone2normal
Phone number normalizer program written in Go

## Info 
phone2normal is a toy program I wrote to get a bit of practice reading/writing from a postgreSQL database. 
This program in its current state opens a connection to the DB, uses the ```resetTable()``` func to remove the existing rows from the table. 
This is so each time you run the program you dont get the same records being inserted. I then use ```Query()``` from the sql/database package to iterate the rows and store those rows in a slice of strings.
Once we have the records stored in a slice, we normalize the phone numbers by removing all non-digit characters and then update the rows with the normalized values.
If the record being updated already exists we simply delete the duplicate row.

### Follow ups
-could explore user of an ORM



