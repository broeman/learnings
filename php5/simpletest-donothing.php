<?php

require_once('simpletest/unit_tester.php');
require_once('simpletest/reporter.php');

class TestOfMysqlTransaction extends UnitTestCase {
}

$test = new TestOfMysqlTransaction();
$test->run(new HtmlReporter());

?>