import numpy as np
from stdform import StdForm


def main():
    interactive = False
    if interactive:
        print("Coefficients of objective function: ", end="")
    c = np.array([float(i) for i in input().split()])
    num_var = c.size  # number of decision variables
    if interactive:
        print("Values of constraints: ", end="")
    b = np.array([float(i) for i in input().split()])
    num_eq = b.size  # number of equations/constraints
    a_list = []
    for i in range(num_eq):
        if interactive:
            print(f"Coefficient of constraint {i+1}: ", end="")
        a_list.append([float(i) for i in input().split()])
    a = np.matrix(a_list, dtype=float)

    std = StdForm(c, a, b)
    print(std)


if __name__ == "__main__":
    main()
