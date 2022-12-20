<script setup lang="ts">
import { PolarArea } from "vue-chartjs";
import { Chart, registerables } from "chart.js";
import { Language } from "~~/types/language";

interface Props {
  languages: Language[] | undefined;
}
const props = defineProps<Props>();
const langLabels: string[] = [];
const langData: number[] = [];

props.languages?.map((a) => {
  langLabels.push(a.name);
  langData.push(a.ratio);
});

Chart.register(...registerables);
const data = {
  labels: langLabels,
  datasets: [
    {
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
      data: langData,
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
    <PolarArea :data="data" :options="options" height="330" />
  </div>
</template>
