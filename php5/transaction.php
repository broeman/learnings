<?php

class MySqlTransaction {
	private $connection;

	function __construct($host, $user, $password, $db) {
		$this->connection = mysql_connect(
			$host, $user, $password, $db, true);
		mysql_select_db($db, $this->connection);
		$this->throwOnMysqlError();
		//$this->begin();
	}

	function __destruct() {
		if (isset($this->connection)) {
			@mysql_query('rollback', $this->connection);
			@mysql_close($this->connection);
		}
	}

	function select($sql) {
		$result = @mysql_query($sql, $this->connection);
		$this->throwOnMysqlError();
		return new MysqlResult($result);
	}

	function execute($sql) {
		mysql_query($sql, $this->connection);
		$this->throwOnMysqlError();	
	}

	function commit() {
/*		$this->execute('commit');
		mysql_close($this->connection);
		unset($this->connection); */
	}

	private function throwOnMysqlError() {
		if ($error = mysql_error($this->connection)) {
			throw new Exception($error);
		}		
	}

	private function begin() {
		$this->execute(
			'set transaction isolation level serializable');
		$this->execute('begin');
	}
}

class MysqlResult {
	private $result;

	function __construct($result) {
		$this->result = $result;
	}

	function next() {
		return mysql_fetch_assoc($this->result);
	}
}

?>