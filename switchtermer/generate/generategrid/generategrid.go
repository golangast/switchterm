package generategrid

import (
	"strings"

	"github.com/golangast/switchterm/switchtermer/db/domain"
	"github.com/golangast/switchterm/switchtermer/db/grid"
	"github.com/golangast/switchterm/switchtermer/db/handler"
	"github.com/golangast/switchterm/switchtermer/switch/switchselector"
	"github.com/golangast/switchterm/switchtermer/switchutility"
)

func Grid() {
	d, err := domain.GetStringDomains()
	switchutility.Checklogger(err, "getting all domains for grid")

	chosendomain := switchselector.MenuInstuctions(d, 1, "purple", "purple", "Which website are you going to run the database server for?")

	h, err := handler.GetStringHandlers()
	switchutility.Checklogger(err, "getting all handler for grid")

	gridhandler := switchselector.MenuInstuctions(h, 1, "purple", "purple", "Whats Name of the handler you want to use?")

	gridName := switchutility.InputScan("what is the name of the grid?")
	gridLayout := switchutility.InputScan(`
	what grid and widgets?
	* grids: (f=full, h=half, m=middler, q=quarter) 
	* widgets: (l=list, s=scrollspy, p=parallax, fo=form, i=image, a=accordian) 
	* "/" indicates which widget goes to which grid in order (i.e. left half and right half)
	* example: f-i h-a/a q-i/f0/l/i
	
	`)

	//store grid
	g := grid.Grid{Name: gridName, Domain: chosendomain, Handler: gridhandler, GridLayout: gridLayout}
	if ok, _ := g.Exists(g.Name); !ok {
		g.Create()
	}

	var htmlstring string

	words := strings.Fields(gridLayout)

	for _, word := range words {
		w := findwidget(strings.TrimLeft(word, "-"))
		switch word[:1] {
		case "f":
			htmlstring += strings.Replace(findgrid(word[:1]), "<!--#w-->", w, 1)
		case "h":
			htmlstring += strings.Replace(findgrid(word[:1]), "<!--#w-->", findwidget(strings.TrimRight(w, "/")), 1)
			htmlstring += strings.Replace(findgrid(word[:1]), "<!--#w-->", findwidget(strings.TrimLeft(w, "/")), 1)
		}
	}

	if err := switchutility.UpdateText(chosendomain+"/assets/templates/"+gridhandler+"/"+gridhandler+".html", `<!-- #grid -->`, `<!-- #grid -->`, htmlstring+"\n"); err != nil {
		switchutility.Checklogger(err, "trying to update the grid in template file")
	}

}
func findgrid(s string) string {
	for _, l := range s {
		switch string(l) {
		case "f":
			return full
		case "h":
			return half
		case "m":
			return middler
		case "q":
			return quarters
		}
	}
	return ""
}
func findwidget(s string) string {
	for _, l := range s {
		switch string(l) {
		case "l":
			return list
		case "s":
			return scrollspy
		case "p":
			return parallax
		case "fo":
			return form
		case "i":
			return image
		case "3i":
			return threepicture
		}
	}
	return ""
}

var full = `<div class="full"><!--#w--></div>`
var half = `<div class="lhalf"><!--#w--></div><div class="rhalf"><!--#w--></div>`
var quarters = `<div class="quarter1"><!--#w--></div>
<div class="quarter2"><!--#w--></div>
<div class="quarter3"><!--#w--></div>
<div class="quarter4"><!--#w--></div>`

var middler = `<div class="middler"><!--#w--></div><div class="middler"><!--#w--></div>`

var image = `
<img loading="lazy" class="op" {{ .nonce }} src="img/gocode.webp" >
`
var threepicture = `
<div class="row">
  <div class="col">
    <div data-bs-spy="scroll" data-bs-target="#list-example" data-bs-smooth-scroll="true"  tabindex="0">
      <div class="row">
        <div class="col ">
          <div class="card col s12 m12 l4">
            <div class="card-image">
              <img  loading="lazy" class="brand-logo d-block w-100 para" width="693"  height="924" {{ .nonce }} src="img/me.webp" alt="logo"></a>
              <h4 id="list-item-1">
              <span class="card-title">About Me</span>
              </h4>
            </div>
            <div class="card-content">
              <p> Going out for walks</p>
            </div>
            <div class="card-action">
           
            </div>
          </div>
          <div class="card col s12 m12 l4">
            <div class="card-image">
              <img  loading="lazy" nonce="{{ .nonce }}" class="brand-logo d-block w-100 op"  {{ .nonce }} src="img/pizza.webp" alt="logo"></a>

              <span class="card-title">Making Pizza</span>
            </div>
            <div class="card-content">
              <p> Love Pizza</p>
            </div>
            <div class="card-action">
            </div>
          </div>
          <div class="card col s12 m12 l4">
            <div class="card-image">
              <img  loading="lazy" class="brand-logo d-block w-100 op" width="100%" height="auto"  {{ .nonce }} src="img/abe.jpg" alt="logo"></a>
              <span class="card-title">Abraham Lincoln</span>
            </div>
            <div class="card-content">
              <p> Love reading about Abraham Lincoln</p>
            </div>
            <div class="card-action">
            </div>
          </div>
        </div>
      </div>

`

var form = `
<form class="logob  center mcenter" action="/userinput" method="POST">
<label class="field " for="email">Email</label><br>
<input class="field " type="text" id="email" name="email" value=""><br>
<label class="field " for="language">What programming language do you like?</label><br>
<input class="field " type="text" id="language" name="language" value="c#"><br>
<label class="field " for="comment">Ask a question</label><br>
<textarea class="field " rows="4" cols="50" name="comment" type="text">Enter text here...</textarea><br>
<input class="hide" type="text" id="sitetoken" name="sitetoken" value="">
<input type="submit" value="Submit">
</form>
`

var list = `

<table class="table table-striped-columns table-dark">
<thead>
  <tr>
	<th scope="col">#</th>
	<th scope="col">Type</th>
	<th scope="col">Why</th>
  
  </tr>
</thead>
<tbody>
  <tr>
	<th scope="row">1</th>
	<td><a href="https://docs.google.com/document/d/1Zb9GCWPKeEJ4Dyn2TkT-O3wJ8AFc-IMxZzTugNCjr-8/edit?usp=sharing" target="_blank">Google Doc Resources</a></td>
	<td>I needed a way as I was learning and building to reference material.</td>
  </tr>

  <tr>
	<th scope="row">2</th>
	<td><a href="https://www.youtube.com/watch?v=HJHCndEVoiA&list=PL_sE11fwtBT-0GqVHEX-tYTBzAIGHelQ6" target="_blank">Youtube</a></td>
	<td>Wanted to demonstrate things that I have built or learned so that others can use them.</td>
	
  </tr>
  <tr>
	<th scope="row">3</th>
	<td><a href="https://medium.com/@snippet22/errors-in-go-1ebfa1c1b883" target="_blank">Medium</a></td>
	<td>Was a way to express ideas in a blog about Go.</td>
	
  </tr>
  <tr>
	<th scope="row">4</th>
	<td><a href="https://twitter.com/ZacharyEndrulat" target="_blank">Twitter</a></td>
	<td>I use this to talk to others and share things.</td>
	
  </tr>
 
</tbody>
</table>


`
var scrollspy = `
<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-QWTKZyjpPEjISv5WaRU9OFeRpok6YctnYmDr5pNlyT2bRjXh0JMhjY6hW+ALEwIH" crossorigin="anonymous">
<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/materialize/1.0.0/css/materialize.min.css">

<div class="row">
<div class="col-3 sidebar sidem">
	<p>This is what I used to build this site.</p>
	<div id="list-example" class="list-group">
		<a class="list-group-item list-group-item-action" href="#video"> - Video on Goservershell</a>
		<a class="list-group-item list-group-item-action" href="#what"> - What does it do?</a>
		<a class="list-group-item list-group-item-action" href="#why"> - Why build this project?</a>
		<a class="list-group-item list-group-item-action" href="#materialize"> - Materialize CSS</a>
		<a class="list-group-item list-group-item-action" href="#bootstrap"> - Bootstrap</a>
		<a class="list-group-item list-group-item-action" href="#go"> - Go</a>
		<a class="list-group-item list-group-item-action" href="#fontawesome"> - Font Awesome</a>

	</div>
</div>
<div class="col">
	<div data-bs-spy="scroll" data-bs-target="#video" data-bs-smooth-scroll="true"
	class="scrollspy-example" tabindex="0">
	<div id="video" class="card">
		<div class="card-header">
		  
		</div>
		<div class="card-body">
			<p class="card-text">
				<ul class="collection">
					<li class="collection-item">
						<div style="height: 0px; overflow: hidden; padding-top: 56.25%; position: relative; width: 100%;">
							<iframe
								style="position: absolute; top: 0px; left: 0px; width: 100%; height: 100%;"
								src="https://tube.rvere.com/embed?v=HJHCndEVoiA&start=1&end=1020"
								title="YouTube video player"
								frameBorder="0"
								allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
								allowFullScreen>
							</iframe>
						</div>
					</li>

					<li class="collection-item">
						<div style="height: 0px; overflow: hidden; padding-top: 56.25%; position: relative; width: 100%;">
							<iframe
								style="position: absolute; top: 0px; left: 0px; width: 100%; height: 100%;"
								src="https://tube.rvere.com/embed?v=EbQZhMHv9oU&start=1&end=700"
								title="YouTube video player"
								frameBorder="0"
								allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
								allowFullScreen>
							</iframe>
						</div>
				</li>
			</p>
		</div>
	</div>
</div>
<div data-bs-spy="scroll" data-bs-target="#what" data-bs-smooth-scroll="true"
class="scrollspy-example" tabindex="0">
<div id="what" class="card">
	<div class="card-header">
	  What does it do?
	</div>
	<div class="card-body">
		<p class="card-text">
			<ul class="collection">
			   
				<li class="collection-item">
					I wanted to make this project because I felt there needed to be a 
					way to generate handlers and routes and be able to update the 
					frontend quickly.  
				</li>
				<li class="collection-item">
				   1. It is a shell for a server.
				</li>
				<li class="collection-item">
					2. It generates routes and handlers that I built from scratch.
				</li>
				<li class="collection-item">
					3. It hot reloads the site every time the file changes.
				</li>
				<li class="collection-item">
					4. It does security that I did use some libraries for.
					   <ul>
						<li><i class="fa-solid fa-chevron-right fa-xs iconsize"></i><span class="textcenter">a. Certificate generating for the server.</span></li>
					   <li> <i class="fa-solid fa-chevron-right fa-xs iconsize"></i><span class="textcenter"></span>b. naunce security for assets and css that I built from scratch.</span></li>
					   <li> <i class="fa-solid fa-chevron-right fa-xs iconsize"></i><span class="textcenter"></span>c. TLS and HTTPS security.</span></li>
					   <li> <i class="fa-solid fa-chevron-right fa-xs iconsize"></i><span class="textcenter"></span>d. It also does jwt security that I did use a library for the methods</span></li>
					   <li> <i class="fa-solid fa-chevron-right fa-xs iconsize"></i><span class="textcenter"></span>e. Encryption for the keys for jwt I had built and used core libraries for.</span></li>
					</ul>
				</li>
				<li class="collection-item">
					4. It also minifies your css, js, and images.  
					It is a mix of libraries and my own coding.
				</li>
			  </ul>
		</p>
	</div>
</div>
</div>
<div data-bs-spy="scroll" data-bs-target="#why" data-bs-smooth-scroll="true"
		class="scrollspy-example" tabindex="0">
		<div id="why" class="card">
			<div class="card-header">
			   What were my motivations?
			</div>
			<div class="card-body">
				<h5 class="card-title"> motivations</h5>
				<p class="card-text">
					I had built a prior project called <a href="https://github.com/golangast/groundup" target="_blank">Groundup</a> that was my first attempt
					at generating a project. It created a gui to generate apps and allow you
					to start and stop them and generate databases and data. But then someone created 
					the gonew command that pretty much made that a single click to generate a project 
					from a repo.  
				</p>
				<p class="card-text">
				
						I wanted to make this project because I felt there needed to be a 
						way to generate handlers and routes and be able to update the 
						frontend quickly.  <a href="https://go.dev/blog/gonew" target="_blank">Gonew command</a>
						So I decided to work on a shell instead that is focused on upkeep.
				</p>
			</div>
		</div>
	</div>
	<div data-bs-spy="scroll" data-bs-target="#materialize" data-bs-smooth-scroll="true"
		class="scrollspy-example" tabindex="0">
		<div id="materialize" class="card">
			<div class="card-header">
				Materialize
			</div>
			<div class="card-body">
				<h5 class="card-title"> Materialize CSS</h5>
				<p class="card-text">I like Materialize because it lets you create pretty components without
					having to mess much with details.
					It also has a nice nav bar for mobile.</p>
				<a href="https://materializecss.com/" class="btn btn-primary" target="_blank">Site</a>
			</div>
		</div>
	</div>
	<div data-bs-spy="scroll" data-bs-target="#bootstrap" data-bs-smooth-scroll="true" class="scrollspy-example"
		tabindex="0">
		<div id="Bootstrap" class="card">
			<div class="card-header">
				Bootstrap
			</div>
			<div class="card-body">
				<h5 class="card-title"> Bootstrap</h5>
				<p class="card-text">Bootstrap is nice for the same reasons as Materialize. It allows you to
					easily customize the look and feel of the application.
				</p>
				<a href="https://getbootstrap.com/" class="btn btn-primary" target="_blank">Site</a>
			</div>
		</div>
	</div>
	<div data-bs-spy="scroll" data-bs-target="#go" data-bs-smooth-scroll="true" class="scrollspy-example"
	tabindex="0">
	<div id="go" class="card">
		<div class="card-header">
			Go
		</div>
		<div class="card-body">
			<h5 class="card-title"> Go</h5>
			<p class="card-text">

			</p>
			<a href="https://go.dev/" class="btn btn-primary" target="_blank">Site</a>
		</div>
	</div>
</div>
<div data-bs-spy="scroll" data-bs-target="#fontawesome" data-bs-smooth-scroll="true" class="scrollspy-example"
tabindex="0">
<div id="fontawesome" class="card">
	<div class="card-header">
		Font Awesome
	</div>
	<div class="card-body">
		<h5 class="card-title"> Font Awesome</h5>
		<p class="card-text">

		</p>
		<a href="https://fontawesome.com" class="btn btn-primary" target="_blank">Site</a>
	</div>
</div>
</div>
</div>
</div>
`
var parallax = `
<div class="parallax-container paratoper para">
    <div class="parallax"><img  rel="prefetch" class="op" src="img/fl.webp" {{ .nonce }} alt="image"></div>
  </div>

`
