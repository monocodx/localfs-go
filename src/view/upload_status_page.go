// Copyright 2024 The localFS Authors.
// Use of this source code is governed by a GPL
// license that can be found in the LICENSE file.

package view

type UploadStatusPageViewModel struct {
	Error     bool
	Message   string
	Filename  string
	Size      string
	Sha256sum string
	NavBar    NavBar
}

const UploadStatusPageTmpl string = `<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <meta http-equiv="X-UA-Compatible" content="ie=edge">
  <title>localFS</title>
  <style>
    @media only screen and (max-width: 480px) {
      body {
        width: 86% !important;
        padding: .85rem !important;
      }
      div.info {
        padding: 1rem !important;
      }
    }
    body {
      margin: auto;
      width: 60%;
      padding: 1.5rem;
      font-weight: 400;
      font-size: 1rem;
      line-height: 1rem;
      font-family: sans-serif;
    }
    div.navbar {
      display: block;
      margin-bottom: 1.5rem;
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
    div.info {
      display: block;
      border-radius: .75rem;
      padding: 1.5rem;
      background-color: #eceff1;
      margin: 1rem 0rem;
    }
    p.info {
      color: #607D8B;
      font-size: 1rem;
      margin-block-start: 0rem;
      margin-block-end: 0rem;
      padding: .5rem;
      line-break: anywhere;
    }
    span.status{
      display: inline-block;
      border-radius: .75rem;
      padding: .5rem .85rem .3rem .75rem;
      font-weight: 500;
      font-size: .94rem;
      margin-bottom: .75rem;
    }
    span.status.success {
      background-color: #69f0ae;
      color: #1b5e20;
    }
    span.status.error {
      background-color: #ffcdd2;
      color: #b71c1c;
    }
    span.message {
      color: #b71c1c;
      font-size: .85rem;
      margin-left: .25rem;
    }
    i.fa-error::before {
      /* Font Awesome Free 6.7.2 by @fontawesome - https://fontawesome.com License - https://fontawesome.com/license/free Copyright 2025 Fonticons, Inc. */
      content: url('data:image/svg+xml;utf8,<svg viewBox="0 0 48 48" xmlns="http://www.w3.org/2000/svg"><path d="m23.999 2.9988c1.3313 0 2.5596 0.70318 3.2346 1.8564l20.251 34.502c0.68442 1.1626 0.68442 2.5971 0.01875 3.7596-0.66567 1.1626-1.9126 1.8845-3.2534 1.8845h-40.503c-1.3407 0-2.5877-0.72193-3.2534-1.8845-0.66567-1.1626-0.6563-2.6064 0.018752-3.7596l20.251-34.502c0.67505-1.1532 1.9033-1.8564 3.2346-1.8564zm0 12.001c-1.247 0-2.2502 1.0032-2.2502 2.2502v10.501c0 1.247 1.0032 2.2502 2.2502 2.2502s2.2502-1.0032 2.2502-2.2502v-10.501c0-1.247-1.0032-2.2502-2.2502-2.2502zm3.0002 21.002a3.0002 3.0002 0 1 0-6.0004 0 3.0002 3.0002 0 1 0 6.0004 0z" fill="%23b71c1c" stroke-width=".093757"/></svg>');
    }    
    i.fa-success::before {
      /* Font Awesome Free 6.7.2 by @fontawesome - https://fontawesome.com License - https://fontawesome.com/license/free Copyright 2025 Fonticons, Inc. */
      content: url('data:image/svg+xml;utf8,<svg viewBox="0 0 48 48" xmlns="http://www.w3.org/2000/svg"><path d="m24 48a24 24 0 1 0 0-48 24 24 0 1 0 0 48zm10.594-28.406-12 12c-0.88125 0.88125-2.3062 0.88125-3.1781 0l-6-6c-0.88125-0.88125-0.88125-2.3062 0-3.1781s2.3062-0.88125 3.1781 0l4.4062 4.4062 10.406-10.416c0.88125-0.88125 2.3062-0.88125 3.1781 0s0.88125 2.3062 0 3.1781z" fill="%231b5e20" stroke-width=".09375"/></svg>');
    }    
    i.fa-error, i.fa-success {
      width: 18px;
      margin-right: .5rem;
      vertical-align: middle;
    }
    i {
      display: inline-block;
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
  <div class="info">
    {{if .Error}}
      <span class="status error">Error</span>
      <span class="message"><i class="fa-error"></i>{{.Message}}</span>
    {{else}}
      <span class="status success"><i class="fa-success"></i>Completed</span>
    {{end}}
    <p class="info">file: {{.Filename}}</p>
    <p class="info">size: {{.Size}}</p>
    <p class="info">hash: {{.Sha256sum}}</p>
  </div>
</body>
</html>
`
