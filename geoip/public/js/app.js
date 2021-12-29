$('#search').on('click', () => {
    ip = $('#ip').val()
    $('#result').val('')
    $.getJSON('/geo/' + ip, (res) => {
        console.log(ip, res)
        $('#result').val(JSON.stringify(res, null, 4))
    })
})