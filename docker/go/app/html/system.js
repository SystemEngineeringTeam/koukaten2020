document.getElementById("submit").onclick = function () {
    var x = document.getElementById("formText").value;
    var y = document.getElementById("formWho").value;
    var ymd = document.getElementById("YMD").value;
    var hour = document.getElementById("hour").value;
    var minute = document.getElementById("forminute").value;
    var s = "-";
    var zi = "時";
    var hun = "分";
    document.getElementById("output").innerHTML =
        ymd + s + hour + zi + minute + hun + x + s + y;
};
