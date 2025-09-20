from simplex.stdform import StdForm

import numpy as np


class SlackForm(StdForm):
    B: np.matrix  # basic variables
    N: np.matrix  # non-basic variables
    v: float  # objective function constant

    def __init__(
        self,
        c: np.matrix,
        a: np.matrix,
        b: np.matrix,
        B: np.matrix,
        N: np.matrix,
        v: float,
    ) -> None:
        super().__init__(c, a, b)
        self.B = B
        self.N = N
        self.v = v

    @staticmethod
    def fromStd(std: StdForm) -> 'SlackForm':
        num_eq = std.b.size
        num_vars = std.c.size

        c_list = []
        for i in range(num_vars):
            c_list.append(std.c[i])
        for i in range(num_eq):
            c_list.append
        c = np.matrix(num_eq+num_vars)
        return SlackForm()
