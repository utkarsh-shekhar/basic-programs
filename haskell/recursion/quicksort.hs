module Quicksort where

quicksort :: Ord a => [a] -> [a]
quicksort []     = []
quicksort [x]    = [x]
quicksort (x:xs) = (quicksort lt) ++ [x] ++ (quicksort gt)
  where
    lt = filter (<= x) xs
    gt = filter (>  x) xs
