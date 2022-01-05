package main

var homeHTML string = `
<!DOCTYPE html>
<html>
<head>
  <link rel="stylesheet" href="./web/style.css">
  <title>Kobo WiFi Transfer</title>
</head>
<body>
<form enctype="multipart/form-data" action="/upload" method="post">
  <input type="file" id="upload-file" class="upload-file" onClick="selectFile()"/>
  </br>
  <p>
    <input type="checkbox" id="upload-converted" class="upload-converted" checked="checked" />
    <label class="converted-level">Convert EPUB to KEPUB</label>
  </p>
  <input type="button" value="Click to upload file" id="upload-btn" class="upload-btn" onClick="uploadFile()"/>
  <div id="upload-result" class="upload-result"></div>
</form>
</body>
<script src="./web/jquery.min.js"></script>
<script src="./web/upload.js"></script>
</html>
`