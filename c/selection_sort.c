//Selection Sort Using C

#include<stdio.h>

int main() {
  int i, j, r, temp, num[25];
  printf("Enter the interval: ");
  scanf("%d", & r);
  printf("Enter %d Values: \n", r);
  for (i = 0; i < r; i++) {
    scanf("%d", & num[i]);
  }
  for (i = 0; i < r; i++) {
    for (j = i + 1; j < r; j++) {
      if (num[i] > num[j]) {
        temp = num[i];
        num[i] = num[j];
        num[j] = temp;
      }
    }
  }

  printf("Sorted Values: ");
  for (i = 0; i < r; i++) {
    printf("%d\t", num[i]);
  }
}
