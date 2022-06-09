$.ajax({
    url: '/api/getReport',
    success: function(data){
     console.log(data);
    let dashChartBarsCnt = jQuery('.js-chartjs-bars')[0]
    let dashChartLinesCnt = jQuery('.js-chartjs-lines')[0]

    let dashChartBarsData = {
        labels: ['周一', '周二', '周三', '周四', '周五', '周六', '周日'],
        datasets: [
            {
                label: '被举报人数',
                borderWidth: 1,
                borderColor: 'rgba(0,0,0,0)',
                backgroundColor: 'rgba(51,202,185,0.5)',
                hoverBackgroundColor: "rgba(51,202,185,0.7)",
                hoverBorderColor: "rgba(0,0,0,0)",
                data: data.a
            }
        ]
    };
    let dashChartLinesData = {
        labels: data.bDate,
        datasets: [
            {
                label: '被举报人数',
                data: data.b,
                borderColor: '#358ed7',
                backgroundColor: 'rgba(53, 142, 215, 0.175)',
                borderWidth: 1,
                fill: false,
                lineTension: 0.5
            }
        ]
    };

    new Chart(dashChartBarsCnt, {
        type: 'bar',
        data: dashChartBarsData
    });

    let myLineChart = new Chart(dashChartLinesCnt, {
        type: 'line',
        data: dashChartLinesData,
    });

}
  
});
