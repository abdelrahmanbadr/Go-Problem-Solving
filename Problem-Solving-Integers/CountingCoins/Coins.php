<?php

class Coins
{
    public function coins($array, $count)
    {
        $result = [];
        foreach ($array as $row) {
            if ($count == 0) {
                break;
            }
            while ($count >= $row) {
                $count = $count - $row;
                $result[] = $row;
            }
        }
        return $result;
    }
}