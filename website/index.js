

$(document).ready(function() {

    console.log(randomTimeOut);
    // ON PAGE LOAD
    typeText(text, index);
});


$("#submitButton").click(function() {

    // values from form
    let link = $("#link").val();
    let key = $("#key").val();
    
    // display values to be sent to backend on the console
    console.log(`SENT VALUE: ${link}`);
    console.log(`SENT KEY: ${key}`);
    
    // POST method for now, can be later changed to JSON if needed
    // URL to be updated
    $.post( "http://localhost:3000/go-server-endpoint", {
        link: link,
        key: key,
        }, (response) => {

                // response is JSON object return from the server
                console.log(response);

                // converting JSON into JavaScript object
                let data = JSON.parse(response);

                // display response in console / on website
                // console.log(data.message);
                // $("#response").html(response);
            }).fail( (jqXHR, textStatus, errorThrown) => {

            console.log('ERROR HADNLING');
            console.log("jqXHR: " + jqXHR);
            console.log("textStatus: " + textStatus);
            console.log("errorThrown: " + errorThrown);
        })
});

let text = "Ultra Secrete Note Pad"
let index=0;
randomTimeOut = Math.floor(Math.random() * 300);

const typeText = (text, index) => {
    
    if (index < text.length) {
         $("#title").text($("#title").text() + text.charAt(index));
         index++;

         setTimeout( () => {
            typeText(text, index);

         }, 110);
    }
}
