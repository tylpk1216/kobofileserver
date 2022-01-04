let $inputName = $('#upload-file');
let $converted = $('#upload-converted');
let $uploadBtn = $('#upload-btn');
let $uploadResult = $('#upload-result');

function selectFile()
{
  $uploadResult.text('');
}

function disableUI(status)
{
  $inputName.prop('disabled', status);
  $uploadBtn.prop('disabled', status);
}

function uploadFile()
{
  if ($inputName[0].files.length == 0) {
    $uploadResult.text('Error: Please select file.');
    return;
  }

  $uploadResult.text('processing...');
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
    url: '/upload',
    type: 'POST',
    timeout: 3600000,
    data : formData,
    processData: false,
    contentType: false,
    success: function(data, result) {
      disableUI(false);
      $uploadResult.text(data);
    },
    error: function(xhr, textStatus, message) {
      disableUI(false);
      $uploadResult.text(message);
    }
  });
}