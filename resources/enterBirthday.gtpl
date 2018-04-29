<!doctype html>
<html>
    <head>
    <title>Enter the Birthdate</title>

    <!-- Required meta tags -->
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">

    <!-- Bootstrap CSS -->
    <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/css/bootstrap.min.css" integrity="sha384-Gn5384xqQ1aoWXA+058RXPxPg6fy4IWvTNh0E263XmFcJlSAwiGgFAW/dAiS6JXm" crossorigin="anonymous">
    </head>
    <body>
    <!--    <form action="/birthdays/save" method="post">
            Name:<input type="text" name="name">
            Birthday:<input type="date" name="birthday">
            <input type="submit" value="Speichern">
        </form>
     -->   
        <form action="/birthdays/save" method="post">
            <div class="form-group">
                <label for="name">Name</label>
                <input type="text" class="form-control" name = "name" id="name" aria-describedby="nameHelp" placeholder="Gebe hier den Namen ein">
                <small id="nameHelp" class="form-text text-muted">Vorname Nachname</small>
            </div>
            <div class="form-group">
                <label for="birthday">Geburtstag</label>
                <input type="date" class="form-control" name="birthday" id="birthday" placeholder="Geburtstag">
            </div>
            <input type="submit" class="btn btn-primary" value="Speichern">
        </form>




        <script src="https://code.jquery.com/jquery-3.2.1.slim.min.js" integrity="sha384-KJ3o2DKtIkvYIK3UENzmM7KCkRr/rE9/Qpg6aAZGJwFDMVNA/GpGFF93hXpG5KkN" crossorigin="anonymous"></script>
        <script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.12.9/umd/popper.min.js" integrity="sha384-ApNbgh9B+Y1QKtv3Rn7W3mgPxhU9K/ScQsAP7hUibX39j7fakFPskvXusvfa0b4Q" crossorigin="anonymous"></script>
        <script src="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/js/bootstrap.min.js" integrity="sha384-JZR6Spejh4U02d8jOt6vLEHfe/JQGiRRSQQxSfFWpi1MquVdAyjUar5+76PVCmYl" crossorigin="anonymous"></script>
    </body>
</html>