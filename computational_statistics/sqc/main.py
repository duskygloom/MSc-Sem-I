from matplotlib.figure import Figure
import matplotlib.pyplot as plt
import numpy as np
import math


class Statistics:
    @staticmethod
    def offline_mean(data: list[int | float]) -> float:
        return sum(data) / len(data)

    @staticmethod
    def offline_variance(data: list[int | float]) -> float:
        mean_d = Statistics.offline_mean(data)
        return sum([i * i for i in data]) / len(data) - mean_d**2

    @staticmethod
    def online_mean(new_data: int | float, old_mean: float, old_size: int) -> float:
        if old_size == 0:
            return new_data
        return (old_size * old_mean + new_data) / (old_size + 1)

    @staticmethod
    def online_variance(
        new_data: int | float, old_var: float, old_mean: float, old_size: int
    ) -> tuple[float, float]:
        if old_size == 0:
            return 0, new_data
        new_mean = Statistics.online_mean(new_data, old_mean, old_size)
        first_term = old_size * (old_size + 1) * old_var
        second_term = old_size * (new_data - old_mean) ** 2
        third_term = (old_size + 1) ** 2
        new_variance = (first_term + second_term) / third_term
        return new_variance, new_mean


class SQC:
    data: list[int | float]
    fig: Figure
    axes: np.ndarray

    def __init__(self, data: list[int | float], k: float, *, scale: int = 1) -> None:
        self.data = data
        self.k = k
        self.fig, self.axes = plt.subplots(4, 1)
        self.fig.set_figwidth(11.69 * scale)
        self.fig.set_figheight(8.27 * scale)
        self.fig.subplots_adjust(hspace=0.75)

    def draw_data(self):
        mean_d = Statistics.offline_mean(self.data)
        var_d = Statistics.offline_variance(self.data)
        width = 3 * math.sqrt(var_d)
        self.axes[0].set_title("Global SQC")
        self.axes[0].plot(np.arange(len(self.data)), self.data)
        self.axes[0].plot(
            np.arange(len(self.data)),
            [mean_d] * len(self.data),
            linestyle="dashed",
        )
        self.axes[0].plot(
            np.arange(len(self.data)),
            [mean_d + width] * len(self.data),
            linestyle="dashed",
            color="grey",
        )
        self.axes[0].plot(
            np.arange(len(self.data)),
            [mean_d - width] * len(self.data),
            linestyle="dashed",
            color="grey",
        )

    def draw_online(self):
        mean_d = self.data[0]
        var_d = 0

        mean_list = []
        var_list = []

        for i in range(len(self.data)):
            var_d, mean_d = Statistics.online_variance(self.data[i], var_d, mean_d, i)
            mean_list.append(mean_d)
            var_list.append(var_d)

        self.axes[1].set_title("Online SQC")
        self.axes[1].plot(np.arange(len(self.data)), self.data)
        self.axes[1].plot(
            np.arange(len(self.data)),
            mean_list,
            linestyle="dashed",
        )
        self.axes[1].plot(
            np.arange(len(self.data)),
            [mean_list[i] + 3 * math.sqrt(var_list[i]) for i in range(len(self.data))],
            linestyle="dashed",
            color="grey",
        )
        self.axes[1].plot(
            np.arange(len(self.data)),
            [mean_list[i] - 3 * math.sqrt(var_list[i]) for i in range(len(self.data))],
            linestyle="dashed",
            color="grey",
        )

    def draw_penalized_data(self):
        data_p = self.data.copy()
        mean_d = self.data[0]
        var_d = 0

        mean_list = []
        var_list = []

        for i in range(len(self.data)):
            data_p[i] = data_p[i] * (1 - self.k) + self.k * mean_d
            var_d, mean_d = SQC.online_penalty(self.data[i], var_d, mean_d, i, self.k)
            mean_list.append(mean_d)
            var_list.append(var_d)

        self.axes[2].set_title("Penalized data (k = 0.45)")
        self.axes[2].plot(np.arange(len(self.data)), self.data)
        self.axes[2].plot(np.arange(len(self.data)), data_p)

    @staticmethod
    def online_penalty(
        new_data: int | float, old_var: float, old_mean: float, old_size: int, k: float
    ) -> tuple[float, float]:
        data_p = (1 - k) * new_data + k * old_mean
        return Statistics.online_variance(data_p, old_var, old_mean, old_size)

    def draw_penalized(self):
        data_p = self.data.copy()
        mean_d = self.data[0]
        var_d = 0

        mean_list = []
        var_list = []

        for i in range(len(self.data)):
            data_p[i] = data_p[i] * (1 - self.k) + self.k * mean_d
            var_d, mean_d = SQC.online_penalty(self.data[i], var_d, mean_d, i, self.k)
            mean_list.append(mean_d)
            var_list.append(var_d)

        self.axes[3].set_title("Penalized SQC (k = 0.45)")
        self.axes[3].plot(np.arange(len(self.data)), self.data)
        self.axes[3].plot(np.arange(len(self.data)), data_p)
        self.axes[3].plot(
            np.arange(len(self.data)),
            mean_list,
            linestyle="dashed",
        )
        self.axes[3].plot(
            np.arange(len(self.data)),
            [mean_list[i] + 3 * math.sqrt(var_list[i]) for i in range(len(self.data))],
            linestyle="dashed",
            color="grey",
        )
        self.axes[3].plot(
            np.arange(len(self.data)),
            [mean_list[i] - 3 * math.sqrt(var_list[i]) for i in range(len(self.data))],
            linestyle="dashed",
            color="grey",
        )

    def draw_graph(self, filename: str = "plot.svg"):
        self.draw_data()
        self.draw_online()
        self.draw_penalized_data()
        self.draw_penalized()
        self.fig.savefig(filename)


def main() -> None:
    for i in range(1, 4):
        with open(f"data_{i}.txt", "r") as f:
            data_str = f.read().strip().split()
            data = [int(i) for i in data_str]
            SQC(data, k=0.45, scale=1).draw_graph(f"plot_{i}.png")


if __name__ == "__main__":
    main()
