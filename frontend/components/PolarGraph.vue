<script setup lang="ts">
import {
  Chart as ChartJS,
  RadialLinearScale,
  ArcElement,
  Tooltip,
  Legend,
} from "chart.js";
import { PolarArea } from "vue-chartjs";

ChartJS.register(RadialLinearScale, ArcElement, Tooltip, Legend);
const data = {
  labels: ["Go", "Vue", "TypeScript", "Ruby", "JavaScript"],
  datasets: [
    {
      label: "My Second dataset",
      backgroundColor: "rgba(255,99,132,0.2)",
      // pointBackgroundColor: "rgba(255,99,132,1)",
      pointBorderColor: "#fff",
      pointHoverBackgroundColor: "#fff",
      // pointHoverBorderColor: "rgba(255,99,132,1)",
      data: [28, 48, 40, 19, 96],
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
    Top 5 for Most Used Language
  </div>
  <div>
    <PolarArea :data="data" :options="options" />
  </div>
</template>
