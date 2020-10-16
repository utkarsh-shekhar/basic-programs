//BUBBLE SORT with Time Complexity : n

#include<stdio.h>

int main() {
  int i, j, r, temp, num[100000];
  printf("Enter the interval: ");
  scanf("%d", & r);
  printf("Enter %d Values \n", r);
  for (i = 0; i < r; i++) {
    scanf("%d", & num[i]);
  }

  for (i = 1; i < r; i++) {
    temp = num[i];
    j = i - 1;
    while (j >= 0 && num[j] > temp) {
      num[j + 1] = num[j];
      j = j - 1;
    }
    num[j + 1] = temp;
  }
  printf("Sorted Values: ");
  for (i = 0; i < r; i++) {
    printf("%d\t", num[i]);
  }
}
