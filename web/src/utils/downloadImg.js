export const downloadImage = (imgsrc, name) => {
  var image = new Image()
  image.setAttribute('crossOrigin', 'anonymous')
  image.onload = function() {
    var canvas = document.createElement('canvas')
    canvas.width = image.width
    canvas.height = image.height
    var context = canvas.getContext('2d')
    context.drawImage(image, 0, 0, image.width, image.height)
    var url = canvas.toDataURL('image/png')

    var a = document.createElement('a')
    var event = new MouseEvent('click')
    a.download = name || 'photo'
    a.href = url
    a.dispatchEvent(event)
  }
  image.src = imgsrc
}
