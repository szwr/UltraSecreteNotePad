let role = 1;
let text = "USNP";
let index = 0;
let textLoc = $("#title");
let speed = 150;
let randomTimeOut = Math.floor(Math.random() * speed);

$(document).ready(function() {
    typeText(text, index, textLoc);
    changeRole();
});

// SWITCH BETWEEN MESSAGE AND LINK ROLE
const changeRole = () => {
    let submitButton = $("#submitButton");

    submitButton.off("click");

    // MESSAGE ROLE
    if (role == 1) {
        // show/hide the content for each role and change the parameter for submit
        $(".message").hide();
        $(".link").show();
        submitButton.click(sendMessage);

        role = 0;
        // show role on footer
        $(".currentRole").html('MESSAGE ROLE');
    }
    // LINK ROLE
    else {
        $(".message").show();
        $(".link").hide();
        submitButton.click(sendLink);

        role = 1;
        $(".currentRole").html('LINK ROLE');
    }
};

$("#toggleRoleButton").click(changeRole);

// writing effect, text to write, index, location of the text
const typeText = (text, index, textLoc) => {
    if (index < text.length) {
        textLoc.text(textLoc.text() + text.charAt(index));

        index++;
        // time out is different for each character
        randomTimeOut = Math.floor(Math.random() * speed);

        setTimeout(() => {
            typeText(text, index, textLoc);
        }, randomTimeOut);
    }
};

// ROLE: MESSAGE -> SEND key-value pair: message-message to server, receive link
function sendMessage() {
    // values from form
    let message = $("#message").val();
    let password = $("#encrypt").val();
    let responseLoc = $("#response");
    let returnedErrorLoc= $("#error");


    // POST method for now, can be later changed to JSON if needed
    // URL to be updated
    $.post(
        "http://localhost:3000/add-value", {
            message: message,
            password: password,
        },
        (response) => {
            // $("#response").html(`${response.link}`);
            // $("#error").html(response.error);

            typeText(response.link, index, responseLoc);

        },).fail((jqXHR, textStatus, errorThrown) => {
            console.log("ERROR HADNLING");
            console.log("jqXHR: " + jqXHR);
            console.log("textStatus: " + textStatus);
            console.log("errorThrown: " + errorThrown);

            typeText(textStatus, index, returnedErrorLoc);
        });
}

// ROLE: MESSAGE -> LINK key-value pair: link-link to server, receive message
function sendLink() {
    // values from form
    let link = $("#link").val();
    let password = $("#decrypt").val();

    // POST method for now, can be later changed to JSON if needed
    // URL to be updated
    $.post(
        "http://localhost:3000/read-db", {
            link: link,
            password: password,
        },
        (response) => {
            // $("#response").html(response.message);
            // $("#error").html(response.error);
            typeText(response.message, index, responseLoc);
        },
    ).fail((jqXHR, textStatus, errorThrown) => {
        console.log("ERROR HADNLING");
        console.log("jqXHR: " + jqXHR);
        console.log("textStatus: " + textStatus);
        console.log("errorThrown: " + errorThrown);

        // $("#error").html(response.error);
        typeText(textStatus, index, returnedErrorLoc);
    });
}
