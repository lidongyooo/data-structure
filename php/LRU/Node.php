<?php

class Node {
    public Node $prev;

    public Node $next;

    public int $key;

    public string $value;

    public function __construct(int $key, string $value)
    {
        $this->key = $key;
        $this->value = $value;
    }
}