# Loop Control `break` and `continue`

There are going to be times when you need to skip a single loop or stop the loop from
running altogether. It's possible to do this with variables and if statements, but there is
an easier way.  
The continue keyword stops the execution of the current loop and starts a new loop.
The post loop logic runs, and the loop condition gets evaluated.
The break keyword also stops the execution of the current loop and it stops any new
loops from running.  
Use continue when you want to skip a single item in a collection; for instance, perhaps
it's okay if one of the items in a collection is invalid, but the rest may be okay to process.
Use break when you need to stop processing when there are any errors in the data and
there's no value in processing the rest of the collection.