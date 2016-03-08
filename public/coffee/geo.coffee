validateGeo = (field)->
  (evt) ->
    fieldVal = $("div##{field} input").val()
    if !fieldVal? or fieldVal is '' or parseInt(fieldVal) is 0
      alert("Marca la localización en el mapa")
      evt.preventDefault()
initGeo = ->
  lat = $('div#lat input[type=hidden]').val()
  lng = $('div#lng input[type=hidden]').val()

  location = if lat isnt '' and lng isnt '' then [lat, lng] else ' '

  $('#geocomplete').geocomplete(
    map: '.map_canvas'
    details: 'form'
    markerOptions: draggable: true
    location: location
  ).on('geocode:result', (event, res) ->
    location = res['geometry']['location']
    keys = Object.keys(location)
    $('div#lat input[type=hidden]').val(location[keys[0]])
    $('div#lng input[type=hidden]').val(location[keys[1]])
  ).on('geocode:dragged', (event, latLng) ->
    $('div#lat input[type=hidden]').val(latLng.lat())
    $('div#lng input[type=hidden]').val(latLng.lng())
  )
