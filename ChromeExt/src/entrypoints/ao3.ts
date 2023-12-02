
const url = window.location.href.split('/') // https://archiveofourown.org/works/24062428/chapters/57906862#workskin
if (url[3] == "works")
    handleStoryPage();

function handleStoryPage() {
    const storyLink = url[4] // 24062428
    const updated = document.querySelector("dd.status")?.textContent
    const chapters = document.querySelectorAll("option")
    const chapterLink = chapters[chapters.length - 1]?.value
    const authorTag = document.querySelector("a[rel=author]")
    const authorLink = authorTag instanceof HTMLAnchorElement ? authorTag['href'].split('/').slice(4).join('/') : ""
    const author = authorTag instanceof HTMLAnchorElement ? authorLink.split('/')[2] : "Unnamed"
    console.log(`https://archiveofourown.org/works/${storyLink}/chapters/${chapterLink}`)
    console.log(updated)
    console.log(author)
    const actions = document.querySelector("ul.work.actions")
    const followLi = document.createElement("li")
    followLi.classList.add("subscribe")

    const followBtn = document.createElement("input")
    var followed: boolean = false;
    followBtn.value = followed ? "Unfollow" : "Follow"
    followBtn.type = "submit"
    function followBtnPressed() {
        followed = !followed
        followBtn.value = followed ? "Unfollow" : "Follow"

    }
    followBtn.addEventListener('click', followBtnPressed)
    followLi.appendChild(followBtn)
    actions?.appendChild(followLi);
}





