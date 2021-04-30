<?php

require_once 'DoubleLinked.php';
require_once 'Node.php';

class LRUCache
{
    protected DoubleLinked $cache;

    protected int $capacity;

    protected array $map = [];

    public function __construct(int $capacity = 3)
    {
        $this->capacity = $capacity;
        $this->cache = new DoubleLinked();
    }

    public function put (int $key, string $value)
    {
        if (isset($this->map[$key])) {
            $this->deleteKey($key);
        }

        if ($this->cache->size() === $this->capacity) {
            $this->removeLeastRecently();
        }

        $this->addRecently($key, $value);
    }

    public function get(int $key)
    {
        if (!isset($this->map[$key])) {
            return -1;
        }

        $this->makeRecently($key);
        $node = $this->map[$key];
        return $node->value;
    }

    public function collections()
    {
        $node = $this->cache->tailNode()->prev;

        while ($node !== $this->cache->headNode()) {
            echo "key：{$node->key}   value：{$node->value}".PHP_EOL;
            $node = $node->prev;
        }
    }

    protected function makeRecently(int $key)
    {
        $node = $this->map[$key];
        $this->cache->remove($node);
        $this->cache->addLast($node);
    }

    protected function addRecently(int $key, string $value)
    {
        $node = new Node($key, $value);
        $this->cache->addLast($node);
        $this->map[$key] = $node;
    }

    protected function deleteKey(int $key)
    {
        $node = $this->map[$key];
        $this->cache->remove($node);
        unset($this->map[$key]);
    }

    protected function removeLeastRecently()
    {
        $node = $this->cache->removeFirst();
        unset($this->map[$node->key]);
    }

}