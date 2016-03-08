isRetina = ->
  mediaQuery = "(-webkit-min-device-pixel-ratio: 1.5),\
    (min--moz-device-pixel-ratio: 1.5)\
    (-o-min-device-pixel-ratio: 3/2),\
    (min-resolution: 1.5dppx)"
  if window.devicePixelRatio > 1
    return true
  if window.matchMedia and window.matchMedia(mediaQuery).matches
    return true
  false

retina = ->
  if !isRetina()
    return
  $('img.2x').map (i, image) ->
    path = $(image).attr('src')
    path = path.replace('.png', '@2x.png')
    path = path.replace('.jpg', '@2x.jpg')
    $(image).attr 'src', path

$(document).ready retina

$(document).ready () ->
    $(".mobile-menu").click () ->
      $(".menu").toggle()

visualizeMap = (options) ->
  urlMap = 'https://bottico.cartodb.com/api/v2/viz/5b5a5396-e6e8-11e4-9020-0e853d047bba/viz.json'
  cartodb.createVis 'map', urlMap, options

onSuccess = (response) ->
  options =
    zoom: 12
    center: [response.coords.latitude, response.coords.longitude]
  visualizeMap(options)

if navigator.geolocation
  navigator.geolocation.getCurrentPosition(onSuccess)
else
  alert "Debes habilitar la geolocalización"
  visualizeMap(null)
