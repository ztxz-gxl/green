$.ajax({
    url: '/api/getData',
    success: function (data) {
        console.log(data);
        var ctx1 = document.getElementById('worldwide-sales').getContext("2d");
        var myChart1 = new Chart(ctx1,
            {
                type: "bar",
                data: {
                    labels:  ["周一", "周二", "周三", "周四", "周五", "周六", "周日"],
                    datasets: [
                        {
                            label: "个人代币奖励", data: [15, 30, 55, 65, 60, 80, 95],
                            backgroundColor: "rgba(0, 156, 255, .7)"
                        },
                        {
                            label: "代币奖励平均值", data: [8, 35, 40, 60, 70, 55, 75],
                            backgroundColor: "rgba(0, 156, 255, .5)"
                        }
                    ]
                },
                options: { responsive: true }
            });
        var ctx2 = document.getElementById("salse-revenue").getContext("2d");
        var myChart2 = new Chart(ctx2,
            {
                type: "line",
                data:
                {
                    labels:
                        ["周一", "周二", "周三", "周四", "周五", "周六", "周日"],
                    datasets: [
                        {
                            label: "平均值", data: [15, 30, 55, 45, 70, 65, 85],
                            backgroundColor: "rgba(0, 156, 255, .5)", fill: true
                        },
                        {
                            label: "个人值",
                            data: [99, 135, 170, 130, 190, 180, 270],
                            backgroundColor: "rgba(0, 156, 255, .3)", fill: true
                        }
                    ]
                },
                options: { responsive: true }
            });
    }
})