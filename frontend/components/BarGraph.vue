<script setup lang="ts">
import { Bar } from "vue-chartjs";
import { Chart, registerables } from "chart.js";
import { Commit } from "~/types/commit";
Chart.register(...registerables);

interface Props {
  commits: Commit[] | undefined;
}
const props = defineProps<Props>();
const date: string[] = [];
const count: number[] = [];

props.commits?.map((a) => {
  date.push(a.date);
  count.push(a.count);
});
const data = {
  labels: date,
  datasets: [
    {
      label: "My Second dataset",
      backgroundColor: [
        "rgba(255, 99, 132, 0.2)",
        "rgba(54, 162, 235, 0.2)",
        "rgba(255, 206, 86, 0.2)",
        "rgba(75, 192, 192, 0.2)",
        "rgba(153, 102, 255, 0.2)",
        "rgba(255, 159, 64, 0.2)",
      ],
      borderColor: [
        "rgba(255,99,132,1)",
        "rgba(54, 162, 235, 1)",
        "rgba(255, 206, 86, 1)",
        "rgba(75, 192, 192, 1)",
        "rgba(153, 102, 255, 1)",
        "rgba(255, 159, 64, 1)",
      ],
      borderWidth: 1,

      data: count,
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
    <Bar :data="data" :options="options" width="400" height="300" />
  </div>
</template>
