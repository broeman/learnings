<?php

require_once('simpletest/unit_tester.php');
require_once('simpletest/reporter.php');
require_once('transaction.php');

class TestOfMysqlTransaction extends UnitTestCase {

	// TEST CASES
	function testCanReadSimpleSelect() {
		$transaction = new MysqlTransaction(
			'localhost', 'me', 'secret', 'test');
		$result = $transaction->select('select 1 as one');
		$row = $result->next();
		$this->assertEqual($row['one'], 1);
	}

	function testShouldThrowExceptionOnBadSelectSyntax() {
		$transaction = new MysqlTransaction(
			'localhost', 'me', 'secret', 'test');
		$this->expectException();
		$transaction->select('not valid SQL');
	}

	function testCanWriteRowAndReadItBack() {
		$this->createSchema();
		$transaction = new MysqlTransaction(
			'localhost', 'me', 'secret', 'test');
		$transaction->execute('insert into numbers (n) values (1)');
		$result = $transaction->select('select * from numbers');
		$row = $result->next();
		$this->assertEqual($row['n'], '1');
		$this->dropSchema();
	}

/* this is quite slow
	function testRowConflictBlowsOutTransaction() {
		$this->setUpRow();
		$one = new MysqlTransaction(
			'localhost', 'me', 'secret', 'test');
		$one->execute('update numbers set n = 2 where n = 1');
		$two = new MysqlTransaction(
			'localhost', 'me', 'secret', 'test');
		try {
			$two->execute('update numbers set n = 3 where n = 1');
			$this->fail('Should have thrown');
		} catch (Exception $e) { }
		$this->dropSchema();
	} */

	// HELPER FUNCTIONS
	function setUpRow() {
		$this->createSchema();
		$transaction = new MysqlTransaction(
			'localhost', 'me', 'secret', 'test');
		$transaction->execute('insert into numbers (n) values (1)');
		$transaction->commit();		
	}

	private function createSchema() {
		$transaction = new MysqlTransaction(
			'localhost', 'me', 'secret', 'test');
		$transaction->execute('drop table if exists numbers');
		$transaction->execute(
			'create table numbers (n integer)');
		$transaction->commit();		
	}

	private function dropSchema() {
		$transaction = new MysqlTransaction(
			'localhost', 'me', 'secret', 'test');
		$transaction->execute('drop table if exists numbers');		
	}
}

$test = new TestOfMysqlTransaction();
$test->run(new HtmlReporter());

?>