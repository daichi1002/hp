<script setup lang="ts">
import {
  Chart as ChartJS,
  Tooltip,
  Legend,
  Title,
  BarElement,
  CategoryScale,
  LinearScale,
} from "chart.js";
import { Bar } from "vue-chartjs";

ChartJS.register(
  CategoryScale,
  LinearScale,
  BarElement,
  Title,
  Tooltip,
  Legend
);
const data = {
  labels: [
    "2022-12-11",
    "2022-12-12",
    "2022-12-13",
    "2022-12-14",
    "2022-12-15",
    "2022-12-16",
    "2022-12-17",
  ],
  datasets: [
    {
      label: "My Second dataset",
      backgroundColor: "rgba(255,99,132,0.2)",
      // pointBackgroundColor: "rgba(255,99,132,1)",
      pointBorderColor: "#fff",
      pointHoverBackgroundColor: "#fff",
      // pointHoverBorderColor: "rgba(255,99,132,1)",
      data: [28, 48, 40, 19, 96, 54, 32, 65],
    },
  ],
};

const options = {
  responsive: true,
  maintainAspectRatio: false,
};
</script>
<template>
  <div class="text-center text-xl font-medium py-4">
    1 week of contributions
  </div>
  <div>
    <Bar :data="data" :options="options" />
  </div>
</template>
