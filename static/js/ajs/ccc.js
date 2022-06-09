$.ajax({
    url: '/api/getData',
    success: function(data){
     console.log(data);
    let Texs = []
    let Typs = []
    for (let i = 0;i < data.a.length;i++){
        Texs[i] = data.a[i].Text
        Typs[i] = parseInt(data.a[i].Type)
    }
     let dashChartLinesData = {
        labels: Texs,
        datasets: [
            {
                label: '今日绿色生活榜单前五',
                data:Typs,
                borderColor: '#28a745',
                backgroundColor: 'rgba(53, 142, 215, 0.175)',
                borderWidth: 1,
                fill: false,
                lineTension: 0.5
            }
        ]
    };

var ctx = document.getElementById("pie-graph").getContext("2d");
let myLineChart = new Chart(ctx, {
    type: 'line',
    data: dashChartLinesData,
});

let Text = []
let Type = []
for (let i = 0;i < data.listData2.length;i++){
    Text[i] = data.listData2[i].Text
    Type[i] = parseInt(data.listData2[i].Type)
}

var doughnutData = {
    labels: Text,
    datasets: [
        {
            data: Type,
            backgroundColor: [
                "#2a7f48",
                "#47e36b",
                "#24c95e",
                "#10d63d",
                "#00ff5a"
            ],
            hoverBackgroundColor: [
                "#2a7f48",
                "#47e36b",
                "#24c95e",
                "#10d63d",
                "#00ff5a"
            ]
        }]
};

var ctx = document.getElementById("doughnut-graph").getContext("2d");
var DoughnutChart = new Chart(ctx, {
    type: 'doughnut',
    data: doughnutData,
    responsive: true
});
    }
})