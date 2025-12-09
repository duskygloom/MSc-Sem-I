from math import inf
import numpy as np
from tabulate import tabulate

np.seterr(divide="ignore")


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
        self.cj = np.zeros(i + j, dtype=float)
        for x in range(j):
            self.cj[x] = cj[x]
        # create extended aij
        self.aij = np.zeros((i, i + j), dtype=float)
        for x in range(i):
            for y in range(j):
                self.aij[x, y] = aij[x, y]
        for x in range(i):
            self.aij[x, j + x] = 1
        # basic variables
        self.xb = [i + x for x in range(j)]

    def next(self) -> None:
        crit_value = self.critical_value
        incoming = self.incoming
        outgoing = self.outgoing
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
    def cb(self) -> list:
        return [self.cj[x] for x in self.xb]

    @property
    def zj(self) -> np.ndarray:
        return np.linalg.matmul(np.array(self.cb), self.aij)

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
            row: list[str | int] = [0 for _ in range(4 + self.i + self.j)]
            row[0] = self.cb[x]
            row[1] = f"x{self.xb[x]+1}"
            row[2] = flat_bi[x]
            for y in range(self.i + self.j):
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

    def __str__(self) -> str:
        headers = ["cb", "xb", "b", *[f"x{x+1}" for x in range(self.i + self.j)], "r"]
        return tabulate(self.table, headers=headers, tablefmt="simple_outline")


def solve(i: int, j: int, cj: np.ndarray, aij: np.ndarray, bi: np.ndarray):
    ti = Tableau(i, j, cj, aij, bi)
    print(ti)
    status = "ok"
    while not ti.is_optimal():
        try:
            ti.next()
            print(ti)
        except Exception as e:
            status = e.args[0]
    if status == "ok":
        for index, value in enumerate(ti.xb):
            print(f"x{value+1} = {ti.bi[index, 0]}")
    else:
        print(status)


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

    solve(i, j, cj, aij, bi)


if __name__ == "__main__":
    main()
