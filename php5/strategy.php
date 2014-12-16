<?php

// A Hello World example in the strategy design pattern
// Jesper Brodersen 2014
// Based on the example from the book 'PHP in Action'

// HtmlDocument is a wrapper that needs a strategy with getContents() to work.
class HtmlDocument {
	private $strategy;

	public function __construct($strategy) {
		$this->strategy = $strategy;
	}

	public function getHtml() {
		return "<html><body>".$this->strategy->getContents() .
		"</body></html>";
	}
}

// A strategy interface that fits with the needs of HtmlDocument
interface HtmlContentStrategy {
	public function getContents();
}

// A strategy implementation
class HelloStrategy implements HtmlContentStrategy {
	var $word;
	public function __construct($word) {
		$this->word = $word;
	}

	public function getContents() {
		return "Hello, ".$this->word."!";
	}
}

// create a new instance of a strategy
$greetings = new HelloStrategy('world');
// using the strategy in the HtmlDocument
$show = new HtmlDocument($greetings);
// showing
echo $show->getHtml();

?>