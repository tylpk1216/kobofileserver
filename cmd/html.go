package main

var homeHTML string = `
<!DOCTYPE html>
<html>
<head>
  <link href="./web/style.css" rel="stylesheet">
  <title>Kobo WiFi Transfer</title>
</head>
<body>
  <div class="bg-white px-8 pt-8 pb-10 md:pt-16 md:px-20 md:pb-20 rounded-xl relative w-full max-w-md md:max-w-xl">
    <div class="mb-8 w-fit py-2 px-4 rounded-xl">
      <h1 class="font-bold text-3xl text-gray-500 mb-1">
        KoboTransfer
      </h1>
      <p class="text-gray-500">
        Run it on Kobo device, then use browser to transfer file to device.
      </p>
    </div>
    <div>
      <form enctype="multipart/form-data" action="/upload" method="post">
        <div class="my-10">
          <input
            class="file:bg-blue-50 file:text-blue-600 file:rounded-2xl file:border-none file:cursor-pointer file:mr-2 file:py-4 file:px-8 w-full border rounded-2xl file:text-base text-sm  text-gray-800 cursor-pointer pr-2"
            id="upload-file" type="file" multiple onClick="selectFile()">
          <div class="my-3 ml-2">
            <label class="relative inline-flex items-center cursor-pointer">
              <input id="upload-converted" type="checkbox" value="" class="sr-only peer" checked="checked">
              <div
                class="w-11 h-6 bg-gray-200 rounded-full peer peer-checked:after:translate-x-full peer-checked:after:border-white after:content-[''] after:absolute after:top-0.5 after:left-[2px] after:bg-white after:border-gray-300 after:border after:rounded-full after:h-5 after:w-5 after:transition-all peer-checked:bg-blue-500 transition-colors">
              </div>
              <span class="ml-3 text-sm font-medium text-gray-800">Convert EPUB to KEPUB</span>
            </label>
          </div>
        </div>
        <button
          class=" bg-blue-500 text-white w-full px-4 py-3 rounded-xl font-medium hover:bg-blue-600 transition-colors"
          type="button" id="upload-btn" onClick="uploadFile()">Upload file</button>
        <hr class="h-px my-8 bg-gray-300 border-0">
        <div id="upload-result" class="upload-result"></div>
        <div class="flex flex-col gap-4">
          <div>
            <div class="result-title">Result</div>
            <div id="tb-result" class="result">&nbsp;</div>
          </div>
          <div>
            <div class="result-title">File Name</div>
            <div id="tb-filename" class="result">&nbsp;</div>
          </div>
          <div>
            <div class="result-title">Saved Time</div>
            <div id="tb-savedtime" class="result">&nbsp;</div>
          </div>
          <div>
            <div class="result-title">Converted Time</div>
            <div id="tb-convertedtime" class="result">&nbsp;</div>
          </div>
        </div>
      </form>
    </div>
  </div>
</body>
<script src="./web/jquery.min.js"></script>
<script src="./web/upload.js"></script>

</html>
`