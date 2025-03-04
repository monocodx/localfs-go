// Copyright 2024 The localFS Authors.
// Use of this source code is governed by a GPL
// license that can be found in the LICENSE file.

package view

type UploadPageViewModel struct {
	Build  string
	Files  []string
	NavBar NavBar
}

const UploadPageTmpl string = `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta http-equiv="expires" content="0">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <meta http-equiv="X-UA-Compatible" content="ie=edge">
  <title>localFS</title>
  <style>
    @media only screen and (max-width: 480px) {
      body {
        width: 86% !important;
        padding: .85rem !important;
      }
      div.center {
        padding: 1rem !important;
      }
      p.lead {
       font-size: 1rem !important;
      } 
      /* 
      div.flex-container .flex-left {
        order: 2 !important;
      }  
      div.flex-container .flex-right {
        order: 1 !important;
      } 
      */
      span.download {
        text-align: left !important;
        margin-right: 0rem !important;
      }
      span.index {
        display:none;
      }
    }
    body {
      margin: auto;
      width: 70%;
      padding: 1.5rem;
      font-weight: 400;
      font-size: 1rem;
      line-height: 1rem;
      font-family: sans-serif;
    }
    div.navbar {
      display: block;
      margin-bottom: 1rem;
    }
    ul {
      list-style-type: none;
      margin: 0;
      padding: 0;
    }
    li {
      display: inline;
      font-size: .9rem;
      color: #607d8b;
    }
    li > a {
      color: #607d8b;
    }
    li+::before { 
      content: " / ";
      margin: 0rem .15rem;
    }
    div.build {
      text-align: right;
      color: #78909c;
      font-weight: 500;
      font-size: .75rem;
      margin-bottom: .25rem;
    }
    div.error {
      display: none;
      border-radius: .75rem;
      padding: 1rem 1.25rem;
      background-color: #ffcdd2;
      color: #b71c1c;
      font-weight: 500;
      font-size: .95rem;
      margin-bottom: .75rem;
    }
    div.center {
      display: block;
      border-radius: .75rem;
      padding: 1.5rem 2.5rem;
      background-color: #e2e7ea;
      text-align: center;
      margin-bottom: 2.5rem;
    }
    span.progress {
      color: #006064;
      background-color: #e2e7ea;
      border-color: #28a745;
      padding: .375rem .75rem;
      line-height: 1.5;
      font-weight: 500;
    }
    div.head {
      border-bottom: .125rem solid #90A4AE;
      margin-bottom: .75rem;
    }
    p.lead {
      color: #607d8b;
      font-size: 1.25rem;
      font-weight: 500;
      margin-bottom: .5rem;
    }
    div.flex-container {
      display: flex;
      flex-direction: row;
      color: #607d8b;
      font-size: 1rem;
      line-break: anywhere;
    }
    div.flex-container.even {
      background-color: #eceff1;
    }  
    div.flex-container .flex-left {
      order: 1;
      flex: 92%; 
      line-height: 1.75rem;
      padding: 1rem .75rem;
      text-overflow: ellipsis;
      white-space: nowrap;
      overflow: hidden;
    }  
    div.flex-container .flex-right {
      order: 2;
      flex: 8%; 
      padding: 1rem .25rem;
    } 
    span.index {
      margin-right: .75rem;
    }
    span.download {
      font-size: 1.25rem;
      display: block;
      text-align: right;
      margin-right: .65rem;
    }
    span.download > a {
      text-decoration: none;
    } 
    i.fa-error::before {
      /* Font Awesome Free 6.7.2 by @fontawesome - https://fontawesome.com License - https://fontawesome.com/license/free Copyright 2025 Fonticons, Inc. */
      content: url('data:image/svg+xml;utf8,<svg viewBox="0 0 48 48" xmlns="http://www.w3.org/2000/svg"><path d="m23.999 2.9988c1.3313 0 2.5596 0.70318 3.2346 1.8564l20.251 34.502c0.68442 1.1626 0.68442 2.5971 0.01875 3.7596-0.66567 1.1626-1.9126 1.8845-3.2534 1.8845h-40.503c-1.3407 0-2.5877-0.72193-3.2534-1.8845-0.66567-1.1626-0.6563-2.6064 0.018752-3.7596l20.251-34.502c0.67505-1.1532 1.9033-1.8564 3.2346-1.8564zm0 12.001c-1.247 0-2.2502 1.0032-2.2502 2.2502v10.501c0 1.247 1.0032 2.2502 2.2502 2.2502s2.2502-1.0032 2.2502-2.2502v-10.501c0-1.247-1.0032-2.2502-2.2502-2.2502zm3.0002 21.002a3.0002 3.0002 0 1 0-6.0004 0 3.0002 3.0002 0 1 0 6.0004 0z" fill="%23b71c1c" stroke-width=".093757"/></svg>');
    }  
    i.fa-error {
      width: 18px;
      margin-right: .5rem;
      vertical-align: middle;
    }
    i.fa-download::before {
      /* Font Awesome Free 6.7.2 by @fontawesome - https://fontawesome.com License - https://fontawesome.com/license/free Copyright 2025 Fonticons, Inc. */
      content: url('data:image/svg+xml;utf8,<svg viewBox="0 0 48 48" xmlns="http://www.w3.org/2000/svg"><path d="m6 33c-3.3094 0-6 2.6906-6 6v3c0 3.3094 2.6906 6 6 6h36c3.3094 0 6-2.6906 6-6v-3c0-3.3094-2.6906-6-6-6h-9.5156l-4.2469 4.2469c-2.3438 2.3438-6.1406 2.3438-8.4844 0l-4.2375-4.2469zm34.5 5.25a2.25 2.25 0 1 1 0 4.5 2.25 2.25 0 1 1 0-4.5z" fill="%239fa8da"/><path d="m27 3c0-1.6594-1.3406-3-3-3s-3 1.3406-3 3v22.753l-6.8812-6.8812c-1.1719-1.1719-3.075-1.1719-4.2469 0s-1.1719 3.075 0 4.2469l12 12c1.1719 1.1719 3.075 1.1719 4.2469 0l12-12c1.1719-1.1719 1.1719-3.075 0-4.2469s-3.075-1.1719-4.2469 0l-6.8719 6.8812z" fill="%23607d8b"/></svg>');
    }
    i.fa-download {
      width: 22px;
      vertical-align: middle;
    }
    i {
      display: inline-block;
    }
    input {
      font-size: 1rem;
    }
    input[type="submit"] {
      color: #fff;
      background-color: #28a745;
      border-color: #28a745;
      border: 1px solid transparent;
      padding: .375rem .75rem;
      line-height: 1.2rem;
      border-radius: .75rem;
    }
    input[type="file"] {
      display: block;
      background-color: #fff;
      width: 100%;
      margin-bottom: 2rem;
      border-radius: .75rem;
    }
    input::file-selector-button {
      color: #fff;
      background-color: #0288d1;
      border-color: #007bff;
      border: 1px solid transparent;
      padding: .375rem .75rem;
      line-height: 1.2rem;
    }
    .processlabel {
      display: "none";
    }
    /* https://blog.hubspot.com/website/css-loading-animation */
    .process {
      /* display: inline-block; */
      display: none;
      border: 4px solid #e2e7ea;
      border-radius: 50%;
      border-top: 4px solid #0097a7;
      border-bottom: 4px solid #006064;
      border-right: 4px solid #00838f;
      width: 20px;
      height: 20px;
      animation: spinner 1s linear infinite;
      vertical-align: middle;
      margin-right: .25rem;
      /* margin: 1rem .5rem; */
    }
    @keyframes spinner {
      0% {transform: rotate(0deg);}
      100% {transform: rotate(360deg);}
    }
  </style>
</head>
<body>
  <div class="navbar">
    <ul>
    {{range $idx, $item := .NavBar.NavItem}}
      <li><a href="{{$item.Link}}">{{$item.Name}}</a></li>
    {{end}}
    <li>{{.NavBar.ActiveItem}}</li>
    </ul>
  </div>
  <div class="build">build#{{.Build}}</div>
  <div id="error" class="error"><i class="fa-error"></i></div>
  <div class="center">
    <form id="uform" method="post" enctype="multipart/form-data" action="/upload/file">
      <input id="ufile" type="file" name="file" />
      <span id="uprocess" class="process"></span>
      <span id="uprocesslabel" class="uprocesslabel"</span>
      <input id="usubmit" type="submit" value="Upload File">
    </form>
  </div>
  <div class="head">
    <p class="lead">Uploaded File(s)</p>
  </div>
  <!-- Listing -->
  {{if gt (len .Files) 0}} {{range $idx, $item := .Files}}{{if zebraCss $idx}}
  <div class="flex-container even">
    <div class="flex-left">
      <span class="index">{{index $idx}}.</span>{{$item}}
    </div>
    <div class="flex-right">
      <span class="download"><a href="/download/{{$item}}" download="{{$item}}"><i class="fa-download"></i></a></span>
    </div>
  </div>
  {{else}}
  <div class="flex-container">
    <div class="flex-left">
      <span class="index">{{index $idx}}.</span>{{$item}}
    </div>
    <div class="flex-right">
      <span class="download"><a href="/download/{{$item}}" download="{{$item}}"><i class="fa-download"></i></a></span>
    </div>
  </div>
  {{end}}{{end}}{{end}}
  <script>
    let errmsg = "No file selected. Please choose a file to upload."
    let error = document.getElementById("error");
    let uform = document.getElementById("uform");
    let ufile = document.getElementById("ufile");
    let usubmit = document.getElementById("usubmit");
    let uprocess = document.getElementById("uprocess");
    let uprocesslabel = document.getElementById("uprocesslabel");

    // https://stackoverflow.com/questions/8861181/8861236#8861236
    window.addEventListener("pageshow", () => {
      // handle form data not reset 
      // when navigate using browser back button
      uform.reset();

      // handle android devices back action
      if (this.event.persisted) {
        uprocess.style.display = "none";
        uprocesslabel.style.display = "none";
        location.reload();
      }
    });

    ufile.onchange = function () {
      if (this.files[0]) {
        // reset
        error.style.display = "none"
      }
    }

    uform.addEventListener("submit", (e) => {
      e.preventDefault();
      if (ufile.value === undefined || ufile.value === "") {
        error.style.display = "block"
        error.appendChild(document.createTextNode(errmsg));
        return
      }

      uprocess.style.display = "inline-block";
      uprocesslabel.style.display = "inline-block";
      uprocesslabel.textContent = "Uploading...";
      usubmit.disabled = true;
      usubmit.style.display = "none";
      uform.submit();

      // prevent choose new file during upload in progress
      // can not use disabled due to iOS will not upload  
      // ufile.disabled = true;
      ufile.onclick = function() {
        return false;
      }
    });
  </script>
</body>
</html>
`
