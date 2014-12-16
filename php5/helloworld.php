<?php

// A Hello World example in OOP
// Jesper Brodersen 2014
// Based on the example from the book 'PHP in Action'

// HTML Wrapper class
abstract class HtmlDocument {
	// returns html-wrapped content
	public function getHtml() {
		return "<html><body>".$this->getContent()."</body></html>";
	} 

	abstract public function getContent();
}

// Hello is using HTML to print out a hello
class Hello extends HtmlDocument {
	private $word;

	// using a constructor to allow for many types of hellos
	function __construct($word) {
		$this->word = $word;
	}

	// using inheritance to override a function with actual content 
	function getContent() {
		return "Hello, ".$this->word."!";
	}
}

// making a new instance of Hello as greetings
$greetings = new Hello('world');
// calling getHthml on the Hello instance, which is inherited from HtmlDocument
echo $greetings->getHtml();

?>