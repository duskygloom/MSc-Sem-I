from math import inf
import numpy as np
from tabulate import tabulate

np.seterr(divide="ignore", invalid="ignore")


class Tableau:
    i: int
    j: int
    cj: np.ndarray
    aij: np.ndarray
    bi: np.ndarray
    xb: list

    def __init__(
        self, i: int, j: int, cj: np.ndarray, aij: np.ndarray, bi: np.ndarray
    ) -> None:
        self.i = i
        self.j = j
        self.bi = bi
        # create extended cj
        self.cj = np.zeros(2 * i + j, dtype=float)
        for x in range(j):
            self.cj[x] = cj[x]
        # create extended aij
        self.aij = np.zeros((i, 2 * i + j), dtype=float)
        for x in range(i):
            for y in range(j):
                self.aij[x, y] = aij[x, y]
        for x in range(i):
            self.aij[x, x + j] = -1
            self.aij[x, x + i + j] = 1
        # basic variables
        self.xb = [i + j + x for x in range(i)]

    def next(self) -> None:
        crit_value = self.critical_value
        incoming = self.incoming
        outgoing = self.outgoing
        # eliminate outgoing if it is artificial
        outgoing_var = self.xb[outgoing]
        if outgoing_var > self.j + self.i:
            self.aij = np.hstack(
                (self.aij[:, :outgoing_var], self.aij[:, outgoing_var + 1 :])
            )
            self.cj = np.hstack((self.cj[:outgoing_var], self.cj[outgoing_var + 1 :]))

        # matrix transformations
        self.xb[outgoing] = incoming
        self.aij[outgoing] /= crit_value
        self.bi[outgoing, 0] /= crit_value
        for x in range(self.i):
            if x == outgoing:
                continue
            factor = self.aij[x, incoming]
            self.aij[x] -= factor * self.aij[outgoing]
            self.bi[x, 0] -= factor * self.bi[outgoing, 0]

    @property
    def cb(self) -> np.ndarray:
        return np.array([self.cj[x] for x in self.xb])

    @property
    def zj(self) -> np.ndarray:
        return np.linalg.matmul(self.cb, self.aij)

    @property
    def incoming(self) -> int:
        zj_cj = self.zj - self.cj
        min_index = -1
        min_value = inf
        for index, value in enumerate(zj_cj.flatten()):
            if value < min_value:
                min_value = value
                min_index = index
        return min_index

    @property
    def ratio(self) -> np.ndarray:
        index = self.incoming
        r = self.bi.flatten() / self.aij[:, index]
        return r

    @property
    def outgoing(self) -> int:
        min_ratio = inf
        min_index = -1
        for index, value in enumerate(self.ratio):
            if value < min_ratio and value >= 0:
                min_ratio = value
                min_index = index
        if min_index == -1:
            raise Exception("unbounded")
        return min_index

    def is_optimal(self) -> bool:
        for i in self.zj - self.cj:
            if i < 0:
                return False
        return True

    @property
    def critical_value(self) -> float:
        return self.aij[self.outgoing, self.incoming]

    @property
    def table(self) -> list:
        # first 2 columns for cb and xb
        # 3rd column for bi
        # next columns for aij
        # final column for r
        ti = []
        flat_bi = self.bi.flatten()
        for x in range(self.i):
            row: list[str | int] = [0 for _ in range(4 + len(self.cj))]
            print(x)
            row[0] = self.cb[x]
            row[1] = f"x{self.xb[x]+1}"
            row[2] = flat_bi[x]
            for y in range(len(self.cj)):
                row[y + 3] = self.aij[x, y]
            row[-1] = self.ratio[x]
            ti.append(row)
        # zj
        zj_row = ["", "", "zj", *self.zj, ""]
        ti.append(zj_row)
        # zj - cj
        zj_row = ["", "", "zj - cj", *(self.zj - self.cj), ""]
        ti.append(zj_row)
        return ti

    @property
    def result(self) -> np.ndarray:
        result = [0 for _ in range(2 * self.i + self.j)]
        for index, value in enumerate(self.xb):
            result[value] = self.bi[index, 0]
        return np.array(result, dtype=float)

    @property
    def optimum_value(self) -> float:
        value = np.linalg.matmul(self.cj, self.result)
        if isinstance(value, float):
            return value
        return value.flatten()[0]

    def __str__(self) -> str:
        headers = [
            "cb",
            "xb",
            "b",
            *[f"x{x+1}" for x in range(len(self.cj))],
            "r",
        ]
        print(headers)
        print(np.array(self.table))
        return tabulate(self.table, headers=headers, tablefmt="simple_outline")


def solve(
    i: int,
    j: int,
    cj: np.ndarray,
    aij: np.ndarray,
    bi: np.ndarray,
    *,
    print_tables: bool = False,
) -> Tableau | None:
    ti = Tableau(i, j, cj, aij, bi)
    if print_tables:
        print(ti)
    while not ti.is_optimal():
        # try:
        ti.next()
        print(ti)
    # except Exception as e:
    #     print(e)
    #     return None
    return ti


def main():
    j = int(input("Number of decision variables: "))
    # j = 2
    # print(f"Number of decision variables: {j}")
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
    temp_array = np.array(temp, dtype=float)
    aij = temp_array[:, :j]
    bi = temp_array[:, j:]
    print()

    ti = solve(i, j, cj, aij, bi, print_tables=True)
    if ti == None:
        print("No valid solutions.")
        return
    for index in range(j):
        print(f"x{index+1} = {ti.result[index]}")
    print(f"Optimum value: {ti.optimum_value}")


if __name__ == "__main__":
    main()
