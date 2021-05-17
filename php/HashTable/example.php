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

