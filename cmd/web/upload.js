let $inputName = $('#upload-file');
let $converted = $('#upload-converted');
let $uploadBtn = $('#upload-btn');
let $uploadResult = $('#upload-result');

let $tbResult = $('#tb-result');
let $tbFileName = $('#tb-filename');
let $tbSavedTime = $('#tb-savedtime');
let $tbConvertedTime = $('#tb-convertedtime');

const PROCESSING = 'processing ...';

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

  if (json.Result == '') {
    $tbResult.text('OK');
    $tbFileName.text(json.FileName);
    $tbSavedTime.text(json.SavedTime);
    $tbConvertedTime.text(json.ConvertedTime);
    return;
  }

  $tbResult.text(json.Result);
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

  let fileName = $inputName[0].files[0];
  formData.append('upload-file', fileName);

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