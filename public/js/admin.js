$(function() {
  $('.edit-button').on('click', function() {
    let tr = $(this).parents('tr')
    let id = $(tr).find('td#id').html()
    let name = $(tr).find('td#name').html()
    let email = $(tr).find('td#email').html()
    let created = $(tr).find('td#creted-at').html()
    let updated = $(tr).find('td#updated-at').html()
    let deleted = $(tr).find('td#deleted-at').html()

    let modal = $('#edit-modal')
    modal.find('#id').val(id)
    modal.find('#name').val(name)
    modal.find('#email').val(email)
    modal.find('#created-at').val(created)
    modal.find('#updated-at').val(updated)
    modal.find('#deleted-at').val(deleted)

    let path

    if ($(this).parents('table').hasClass('admins')) {
      path = '/admin/user/' + id
    }
    else {
      path = '/user/' + id
    }

    modal.find('form').attr('action', path)

    $('body').append('<div id="modal-bg"></div>')
    modalResize()
    $('#modal-bg, #edit-modal').fadeIn('normal')
  })

  $('.delete-button').on('click', function() {
    let id = $(this).parents('tr').find('td#id').html()

    $.ajax({
      type: 'DELETE',
      url: '/admin/user/' + id,
      success: function() {
        location.reload()
      },
      error: function(jqXhr) {
        console.log('failed')
      }
    })
  })

  $(document).on('click', '#modal-bg', function() {
    modalClose()
  })

  $(document).on('click', '#edit-modal .save-button', function() {
    $('#edit-modal .error').remove()
    let data = $('form').serializeArray()

    $.ajax({
      type: 'PUT',
      url: $('form').attr('action'),
      data: data,
      success: function() {
        location.reload()
      },
      error: function(jqXhr) {
        $('#edit-modal .error').append(jqXhr.responseText)
      }
    })
  })

  $(document).on('click', '#edit-modal .cancel-button', function() {
    modalClose()
  })

  function modalClose() {
    $('#modal-bg, #edit-modal').fadeOut('normal', function() {
      $('#modal-bg').remove()
      $('#edit-modal .error').remove()
    })
  }

  function modalResize() {
    let w = $(window).width()
    let h = $(window).height()
    let cw = $('#edit-modal').outerWidth()
    let ch = $('#edit-modal').outerHeight()

    $('#edit-modal').css({
      'left': ((w - cw) / 2) + 'px',
      'top': ((h- ch) / 2) + 'px'
    })
  }
})