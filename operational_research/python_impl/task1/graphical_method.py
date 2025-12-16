import numpy as np
from numpy.linalg import LinAlgError


def solve(
    i: int, j: int, cj: np.ndarray, aij: np.ndarray, bi: np.ndarray
) -> np.ndarray:
    # How to solve?
    #
    # First: We find where the lines representing the constraints are
    # intersecting the X-axis. We choose the point with the lower X
    # value as one corner point.
    # Second: We find the where the constraints are intersecting the
    # Y-axis and choose the lower Y value as the second corner.
    # Third: We find the intersection point of the two constraints
    # and choose it as the final corner point.
    #
    # Finally, we choose one of the three points with the least
    # objective function value.

    points = bi / aij

    min_points = np.zeros(shape=(j, j), dtype=float)
    for x in range(j):
        min_points[x, x] = points[:, x : x + 1].min()

    temp = []
    for x in range(i):
        for y in range(i):
            if x == y:
                continue
            try:
                inv = np.linalg.inv(np.array([aij[x], aij[y]]))
            except LinAlgError:
                continue
            r = np.linalg.matmul(inv, np.array([bi[x], bi[y]])).flatten().tolist()
            temp.append(r)
    intersection_points = np.array(temp)

    corners = np.array([*intersection_points, *min_points])
    max_value = 0
    max_corner = np.array([0, 0])

    for corner in corners:
        value = np.linalg.matmul(cj, corner)
        if value > max_value:
            max_value = value
            max_corner = corner

    return max_corner


def main():
    # j = int(input("Number of decision variables: "))
    j = 2
    print(f"Number of decision variables: {j}")
    print()

    print("Objective function")
    t = [int(x) for x in input(f"Enter {j} coefficients: ").split()]
    assert len(t) == j
    cj = np.array(t, dtype=float)
    print()

    i = int(input("Number of constraints: "))
    # i = 2
    # print(f"Number of constraints: {i}")
    print()

    temp = []
    print("Constraints")
    for x in range(i):
        t = [int(x) for x in input(f"Enter constraint {x+1}: ").split()]
        assert len(t) == j + 1
        temp.append(t)
    temp_array = np.array(temp, dtype=float)
    aij = temp_array[:, :j]
    bi = temp_array[:, j:]
    print()

    soln = solve(i, j, cj, aij, bi)
    print(f"x, y = {soln}")
    print(f"Z = {np.linalg.matmul(soln, cj)}")


if __name__ == "__main__":
    main()
