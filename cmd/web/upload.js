let $inputName = $('#upload-file');
let $converted = $('#upload-converted');
let $uploadBtn = $('#upload-btn');
let $uploadResult = $('#upload-result');

let $tbResult = $('#tb-result');
let $tbFileName = $('#filename');
let $tbSavedTime = $('#tb-savedtime');
let $tbConvertedTime = $('#tb-convertedtime');

const PROCESSING = 'processing ...';

const ORIGINAL_CSS = 0;
const SMALL_CSS = 1;

let curCSSFlag = ORIGINAL_CSS;

function selectFile()
{
  $uploadResult.text('');
  $tbResult.text('');
  $tbFileName.text('');
  $tbSavedTime.text('');
  $tbConvertedTime.text('');
}

function disableUI(status)
{
  $inputName.prop('disabled', status);
  $uploadBtn.prop('disabled', status);
}

function newXHR()
{
  let xhr = new window.XMLHttpRequest();
  xhr.upload.addEventListener('progress', function(evt) {
    if (evt.lengthComputable) {
      let percentComplete = Number.parseFloat((evt.loaded / evt.total) * 100).toPrecision(4);
      $uploadResult.text(`(${percentComplete}%) ${PROCESSING}`);
    }
  }, false);
  return xhr;
}

function generateResult(msg)
{
  $uploadResult.text('');

  if (msg.indexOf('{') == -1) {
    $tbResult.text(msg);
    return
  }

  let json = JSON.parse(msg);

  $tbResult.text(json.Result);
  $tbFileName.text(json.FileName);
  $tbSavedTime.text(json.SavedTime);
  $tbConvertedTime.text(json.ConvertedTime);
}

function loadCSSFile(fileName){
  var fileRef = document.createElement('link')
  fileRef.setAttribute('rel', 'stylesheet')
  fileRef.setAttribute('type', 'text/css')
  fileRef.setAttribute('href', fileName)
  document.getElementsByTagName('head')[0].appendChild(fileRef);
}

function removeCSSFile(fileName){
  var allItems = document.getElementsByTagName('link');
  for (var i = 0; i < allItems.length; i++) {
    var item = allItems[i];
    if (!item) continue;
	var attribute = item.getAttribute('href');
    if (!attribute) continue;
	if (attribute.indexOf(fileName) == -1) continue;
    
	item.parentNode.removeChild(allItems[i]);
  }
}

function toggleCSS()
{
  if (curCSSFlag == SMALL_CSS) {
  	removeCSSFile('web/style_small.css');
    loadCSSFile('web/style_ori.css');
    curCSSFlag = ORIGINAL_CSS;
  } else {
  	removeCSSFile('web/style_ori.css');
    loadCSSFile('web/style_small.css');
    curCSSFlag = SMALL_CSS;
  }
}

function uploadFile()
{
  if ($inputName[0].files.length == 0) {
    $uploadResult.text('Error: Please select file.');
    return;
  }

  $uploadResult.text(PROCESSING);
  disableUI(true);

  let formData = new FormData();

  for (let i = 0; i < $inputName[0].files.length; i++) {
    let fileName = $inputName[0].files[i];
    formData.append('upload-file', fileName);
  }

  let converted = '0';
  if ($converted.prop('checked')) {
    converted = '1';
  }
  formData.append('upload-converted', converted);

  $.ajax({
    xhr: newXHR,
    url: '/upload',
    type: 'POST',
    timeout: 3600000,
    data : formData,
    processData: false,
    contentType: false,
    success: function(data, result) {
      disableUI(false);
      generateResult(data);
    },
    error: function(xhr, textStatus, message) {
      disableUI(false);
      generateResult(message);
    }
  });
}
