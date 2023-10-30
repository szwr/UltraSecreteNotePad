

$(document).ready(function() {
   
});


$("#submitButton").click(function() {
    let link = $("#link").val();
    let key = $("#key").val();
    
    console.log(`SENT VALUE: ${link}`);
    console.log(`SENT KEY: ${key}`);
    
    // You can process and display the data here
    let sentData = "Value 1: " + link + "<br>Value 2: " + key;
    // $("#response").html(sentData);

    $.ajax( {
            url:"http://localhost:3000/go-server-endpoint",
            type: "POST",
            data: {
                value1: link,
                value2: key,
            }, success: (response) => {

                    console.log(response)
                    // $("#response").html(response);
                },
                error: (jqXHR, textStatus, errorThrown) => {

                    console.log('ERROR HADNLING');
                    console.log("jqXHR: " + jqXHR);
                    console.log("textStatus: " + textStatus);
                    console.log("errorThrown: " + errorThrown);

                }
            })
});