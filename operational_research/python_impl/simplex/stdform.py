import numpy as np
from matutils import left_pad_mat, concat_mat_horiz, concat_mat_vert


class StdForm:
    c: np.matrix  # coefficients of obj func
    a: np.matrix  # coefficients of constraints
    b: np.matrix  # values of constraints

    def __init__(self, c: np.matrix, a: np.matrix, b: np.matrix) -> None:
        self.c = c
        self.a = a
        self.b = b

    def __repr__(self) -> str:
        return self.__str__()

    def __str__(self) -> str:
        obj_mat = left_pad_mat(self.c, 1)
        eq_mat = concat_mat_horiz(self.b.transpose(), self.a)
        mat = concat_mat_vert(obj_mat, eq_mat)
        return mat.__str__()
