$(function() {
  // Create a socket
  let socket = new WebSocket('ws://' + window.location.host + '/messages/socket')

  // Display a message
  let display = function(event) {
    $('#thread').append(tmpl('message_tmpl', {event: event}))
    $('#thread').scrollTo('max')
  }

  // Message received on the socket
  socket.onmessage = function(event) {
    console.log(JSON.parse(event.data))
    // display(JSON.parse(event.data))
  }

  $('#send').click(function(e) {
    let message = $('#message').val()
    $('#message').val('')
    socket.send(message)
  })

  $('#message').keypress(function(e) {
    if (e.charCode == 13 || e.keyCode == 13) {
      $('#send').click()
      e.preventDefault()
    }
  })
})