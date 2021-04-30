<?php

require_once 'LRUCache.php';

$lru = new LRUCache();

$lru->put(1, 'a');
$lru->put(2, 'b');
$lru->put(3, 'c');
$lru->put(4, 'd');
$lru->put(1, 'a');

$lru->collections();