<?php

class LRUCache {

    protected $capacity,$cache,$map;

    /**
     * @param Integer $capacity
     */
    function __construct($capacity) {
        $this->capacity = $capacity;
        $this->cache = new DoubleLinked();
        $this->map = [];
    }

    /**
     * @param Integer $key
     * @return Integer
     */
    function get($key)
    {
        if (isset($this->map[$key])) {
            $node = $this->map[$key];
            $this->makeRecently($key);
            return $node->value;
        }

        return -1;
    }

    /**
     * @param Integer $key
     * @param Integer $value
     * @return NULL
     */
    function put($key, $value)
    {
        if (isset($this->map[$key])) {
            $node = $this->map[$key];
            $this->cache->remove($node);
        }

        if ($this->capacity === $this->cache->size()) {
            $this->removeLeastRecently();
        }

        $this->addRecently($key, $value);
    }

    function makeRecently($key)
    {
        $node = $this->map[$key];
        $this->cache->remove($node);
        $this->cache->addLast($node);
    }

    function addRecently($key, $value)
    {
        $node = new Node($key, $value);
        $this->map[$key] = $node;
        $this->cache->addLast($node);
    }

    function deleteKey($key)
    {
        $node = $this->map[$key];
        $this->cache->remove($key);
        unset($this->map[$key]);
    }

    function removeLeastRecently()
    {
        $deletedNode = $this->cache->removeLast();
        unset($this->map[$deletedNode->key]);
    }
}

Class DoubleLinked
{
    protected $size,$head,$tail;

    public function __construct ()
    {
        $this->size = 0;
        $this->head = new Node(0, 0);
        $this->tail = new Node(0, 0);
        $this->head->next = $this->tail;
        $this->tail->prev = $this->head;
    }

    public function addLast($node)
    {
        $node->prev = $this->head;
        $node->next = $this->head->next;
        $this->head->next->prev = $node;
        $this->head->next = $node;

        $this->size++;
    }

    public function remove($node)
    {
        $node->prev->next = $node->next;
        $node->next->prev = $node->prev;

        $this->size--;
    }

    public function removeLast()
    {
        if ($this->head->next === $this->tail) {
            return null;
        }

        $lastNode = $this->tail->prev;
        $this->remove($lastNode);

        return $lastNode;
    }

    public function size()
    {
        return $this->size;
    }
}

Class Node
{
    public $key,$value,$prev,$next;

    public function __construct($key, $value)
    {
        $this->key = $key;
        $this->value = $value;
    }
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * $obj = LRUCache($capacity);
 * $ret_1 = $obj->get($key);
 * $obj->put($key, $value);
 */

$lRUCache = new LRUCache(2);
$lRUCache->put(1, 1); // 缓存是 {1=1}
$lRUCache->put(2, 2); // 缓存是 {1=1, 2=2}
$lRUCache->get(1);    // 返回 1
$lRUCache->put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
$lRUCache->get(2);    // 返回 -1 (未找到)
$lRUCache->put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
$lRUCache->get(1);    // 返回 -1 (未找到)
$lRUCache->get(3);    // 返回 3
$lRUCache->get(4);    // 返回 4
