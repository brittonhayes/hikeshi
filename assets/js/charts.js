import Chart from 'chart.js';

const OverviewCtx = document.getElementById('overview-bar');
const PieCtx = document.getElementById('overview-pie');
const OverviewBar = new Chart(OverviewCtx, {
    type: 'line',
    canvas: {
        parentNode: {
            style: {
                width: "20%",
            }
        }
    },
    maintainAspectRatio: false,
    data: {
        labels: ['October', 'November', 'December', 'January', 'February', 'March'],
        datasets: [{
            label: 'Incidents by Month',
            data: [12, 19, 3, 5, 2, 3],
            backgroundColor: 'rgb(27,99,222,0.2)',
            borderColor: 'rgb(19,132,253,0.5)',
            pointBackgroundColor: [
                'rgba(54, 162, 235, 0.2)',
                'rgba(255, 206, 86, 0.2)',
                'rgba(75, 192, 192, 0.2)',
                'rgba(153, 102, 255, 0.2)',
                'rgba(255, 159, 64, 0.2)'
            ],
            pointBorderColor: [
                'rgba(255, 99, 133, 1)',
                'rgba(54, 162, 235, 1)',
                'rgba(255, 206, 86, 1)',
                'rgba(75, 192, 192, 1)',
                'rgba(153, 102, 255, 1)',
                'rgba(255, 159, 64, 1)'
            ],
            borderWidth: 2
        }]
    },
    options: {
        scales: {
            yAxes: [{
                ticks: {
                    beginAtZero: true
                }
            }]
        }
    }
});

const OverviewPie = new Chart(PieCtx, {
    type: 'pie',
    canvas: {
        parentNode: {
            style: {
                width: "20%",
            }
        }
    },
    maintainAspectRatio: true,
    data: {
        labels: ['Red', 'Blue'],
        datasets: [{
            label: 'Percentage of incidents closed',
            data: [12, 19],
            backgroundColor: [
                'rgba(255, 99, 132, 0.2)',
                'rgba(54, 162, 235, 0.2)',
                'rgba(255, 206, 86, 0.2)',
                'rgba(75, 192, 192, 0.2)',
                'rgba(153, 102, 255, 0.2)',
                'rgba(255, 159, 64, 0.2)'
            ],
            borderColor: [
                'rgba(255, 99, 133, 1)',
                'rgba(54, 162, 235, 1)',
                'rgba(255, 206, 86, 1)',
                'rgba(75, 192, 192, 1)',
                'rgba(153, 102, 255, 1)',
                'rgba(255, 159, 64, 1)'
            ],
            borderWidth: 2
        }]
    },
    options: {
        scales: {
            yAxes: [{
                ticks: {
                    beginAtZero: true
                }
            }]
        }
    }
});
