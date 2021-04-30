<?php

require_once 'Node.php';

class DoubleLinked
{

    protected Node $head, $tail;

    protected int $size;

    public function __construct()
    {
        $this->head = new Node(0, 0);
        $this->tail = new Node(0, 0);
        $this->head->next = $this->tail;
        $this->tail->prev = $this->head;
        $this->size = 0;
    }

    public function tailNode()
    {
        return $this->tail;
    }

    public function headNode()
    {
        return $this->head;
    }

    public function size()
    {
        return $this->size;
    }

    public function addLast(Node $node)
    {
        $node->next = $this->tail;
        $node->prev = $this->tail->prev;
        $this->tail->prev->next = $node;
        $this->tail->prev = $node;

        $this->size++;
    }

    public function remove(Node $node)
    {
        $node->prev->next = $node->next;
        $node->next->prev = $node->prev;

        $this->size--;
    }

    public function removeFirst()
    {
        if ($this->head->next === $this->tail) {
            return null;
        }

        $firstNode = $this->head->next;
        $this->remove($firstNode);

        return $firstNode;
    }

}