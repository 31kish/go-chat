$(function() {
  $('.edit-button').on('click', function() {
    console.log("edit button clicked")
    let id = $(this).parents('tr').find('td#id').html()
    console.log(id)
  })

  $('.delete-button').on('click', function() {
    console.log('delete button clicked')
  })
})