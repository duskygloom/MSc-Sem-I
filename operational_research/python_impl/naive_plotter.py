from typing import Tuple
from matplotlib.axes import Axes
from matplotlib.figure import Figure
import numpy as np
from numpy.linalg import LinAlgError
import matplotlib.pyplot as plt


def plot(
    i: int, j: int, cj: np.ndarray, aij: np.ndarray, bi: np.ndarray
) -> Tuple[Figure, Axes]:
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

    fig, axes = plt.subplots()

    # Each point in points is a collection of axes intercepts
    # for a particular constraint.
    # E.g. point[0] and point[1] are the X and Y intercepts
    # respectively.
    points = bi / aij

    # plot constraints
    for point in points:
        axes.plot([0, point[0]], [point[1], 0], linestyle="dashed")

    # minimum intercepts / corner points
    max_points = np.zeros(shape=(j, j), dtype=float)
    min_points = np.zeros(shape=(j, j), dtype=float)
    for x in range(j):
        min_points[x, x] = points[:, x : x + 1].min()
        max_points[x, x] = points[:, x : x + 1].max()

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
            # check if all x, y's are positive
            all_positive = True
            for m in r:
                if m < 0:
                    all_positive = False
            if not all_positive:
                continue
            # check if x, y's satisfy all constraints
            valid_point = True
            for index, constraint in enumerate(aij):
                constraint_value = np.linalg.matmul(constraint, np.array(r))
                if constraint_value > bi[index]:
                    valid_point = False
                    break
            if not valid_point:
                continue
            # append if not already exists
            if r not in temp:
                temp.append(r)
    intersection_points = np.array(temp)

    corners = np.array([[0, 0], *intersection_points, *min_points])
    max_value = 0
    max_corner = np.array([0, 0])

    for corner in corners:
        value = np.linalg.matmul(cj, corner)
        axes.annotate(str(value), corner)
        if value > max_value:
            max_value = value
            max_corner = corner

    # sort all corner points
    centroid = corners.mean(axis=0)
    angles = np.arctan(corners[:, 1] - centroid[1], corners[:, 0] - centroid[0])
    corners_sorted = corners[np.argsort(angles)]
    axes.fill(corners_sorted[:, 0], corners_sorted[:, 1], color="cyan")

    axes.scatter(corners[:, 0], corners[:, 1])
    axes.scatter([max_corner[0]], [max_corner[1]], color="red")

    # plot initial iso-profit line
    slope = -cj[0] / cj[1]
    c = 0
    while True:
        x1, y1 = 0, c
        x2, y2 = -c / slope, 0
        # check if points are satisfying at least 1 constraint
        valid_points = True
        for index, constraint in enumerate(aij):
            value_1 = np.linalg.matmul(constraint, np.array([x1, y1]))
            value_2 = np.linalg.matmul(constraint, np.array([x2, y2]))
            if value_1 > bi[index] and value_2 > bi[index]:
                valid_points = False
                break
        if not valid_points:
            break
        axes.plot([x1, x2], [y1, y2], color="lightgray")
        c += 0.5
    return fig, axes


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
    # print()

    temp = []
    print("Constraints")
    for x in range(i):
        t = [int(x) for x in input(f"Enter constraint {x+1}: ").split()]
        assert len(t) == j + 1
        temp.append(t)
    temp_array = np.array(temp)
    aij = temp_array[:, :j]
    bi = temp_array[:, j:]
    print()

    fig, axes = plot(i, j, cj, aij, bi)
    plt.show()


if __name__ == "__main__":
    main()
