async function fetchAndDisplayPosts() {
  try {
    const response = await fetch("https://idleworkshop.com/api/get-posts");

    if (!response.ok) {
      throw new Error(`Error: ${response.statusText}`);
    }

    const data = await response.json();
    console.log("here is the data", data); // debug

    const posts = data.posts;

    if (!Array.isArray(posts)) {
      throw new Error("Posts is not an array.");
    }

    const postsContainer = document.getElementById("posts-container");
    postsContainer.innerHTML = "";

    for (const post of posts) {
      //container for each post preview: will be displaying one of these for each post.
      const blogPreview = document.createElement("div");
      blogPreview.classList.add("blog-preview");

      //container for the header stuff
      const header = document.createElement("header");

      const title = document.createElement("h2");
      title.textContent = post.title;
      const metadata = document.createElement("div");
      metadata.classList.add("metadata");
      const dateContainer = document.createElement("span");
      dateContainer.classList.add("metadata-item");
      const calIcon = document.createElement("i");
      calIcon.classList.add("fas", "fa-calendar");
      const dateText = document.createElement("span");
      dateText.textContent = ` ${new Date(post.date).toLocaleDateString()}`;
      dateContainer.appendChild(calIcon);
      dateContainer.appendChild(dateText);

      const readTimeContainer = document.createElement("span");
      readTimeContainer.classList.add("metadata-item");
      const clockIcon = document.createElement("i");
      clockIcon.classList.add("fas", "fa-clock");
      const readTimeText = document.createElement("span");
      readTimeText.textContent = ` ${post.readtime}`;
      readTimeContainer.appendChild(clockIcon);
      readTimeContainer.appendChild(readTimeText);

      const topicContainer = document.createElement("span");
      topicContainer.classList.add("metadata-item");
      const topicIcon = document.createElement("i");
      topicIcon.classList.add("fas", "fa-pen-to-square");
      const topicText = document.createElement("span");
      topicText.textContent = ` ${post.topic}`;
      topicContainer.appendChild(topicIcon);
      topicContainer.appendChild(topicText);

      metadata.appendChild(dateContainer);
      metadata.appendChild(readTimeContainer);
      metadata.appendChild(topicContainer);

      header.appendChild(title);
      header.appendChild(metadata);

      // the nutshell sub-heading
      const nutshell = document.createElement("p");
      nutshell.classList.add("nutshell");
      nutshell.textContent = post.nutshell;

      // the 1-2 paragraph summary
      const summary = document.createElement("md-block");
      summary.textContent = post.summary;

      //footer with tags and button
      const footer = document.createElement("footer");
      const tagsWrapper = document.createElement("div");
      tagsWrapper.classList.add("tags-wrapper");

      const tagIcon = document.createElement("i");
      tagIcon.classList.add("fas", "fa-tags");
      tagsWrapper.appendChild(tagIcon);
      const tagList = document.createElement("ul");
      for (const tag of post.tags) {
        const tagElement = document.createElement("li");
        tagElement.textContent = tag;
        tagList.appendChild(tagElement);
      }
      tagsWrapper.appendChild(tagList);

      const button = document.createElement("button");
      button.classList.add("button");
      button.textContent = "Check it out";
      button.onclick = () => displayPost(post);

      footer.appendChild(tagsWrapper);
      footer.appendChild(button);

      /* Save this for the actual blog post page
      const contentElement = document.createElement("md-block");
      contentElement.textContent = post.content;
      */

      blogPreview.appendChild(header);
      blogPreview.appendChild(nutshell);
      blogPreview.appendChild(summary);
      blogPreview.appendChild(footer);

      postsContainer.appendChild(blogPreview);
    }
  } catch (error) {
    console.error(error);
    const postsContainer = document.getElementById("posts-container");
    postsContainer.innerHTML = "<p>Failed to fetch posts.</p>";
  }
}
document.addEventListener("DOMContentLoaded", fetchAndDisplayPosts);

function displayPost(post) {
  const blogPost = document.createElement("div");
  blogPost.classList.add("blog-post");

  const postContainer = document.getElementById("posts-container");
  postContainer.innerHTML = "";

  const header = document.createElement("header");

  const title = document.createElement("h1");
  title.textContent = post.title;
  header.appendChild(title);

  const metadata = document.createElement("div");
  metadata.classList.add("metadata");
  const dateContainer = document.createElement("span");
  dateContainer.classList.add("metadata-item");
  const calIcon = document.createElement("i");
  calIcon.classList.add("fas", "fa-calendar");
  const dateText = document.createElement("span");
  dateText.textContent = ` ${new Date(post.date).toLocaleDateString()}`;
  dateContainer.appendChild(calIcon);
  dateContainer.appendChild(dateText);

  const readTimeContainer = document.createElement("span");
  readTimeContainer.classList.add("metadata-item");
  const clockIcon = document.createElement("i");
  clockIcon.classList.add("fas", "fa-clock");
  const readTimeText = document.createElement("span");
  readTimeText.textContent = ` ${post.readtime}`;
  readTimeContainer.appendChild(clockIcon);
  readTimeContainer.appendChild(readTimeText);

  const topicContainer = document.createElement("span");
  topicContainer.classList.add("metadata-item");
  const topicIcon = document.createElement("i");
  topicIcon.classList.add("fas", "fa-pen-to-square");
  const topicText = document.createElement("span");
  topicText.textContent = ` ${post.topic}`;
  topicContainer.appendChild(topicIcon);
  topicContainer.appendChild(topicText);

  metadata.appendChild(dateContainer);
  metadata.appendChild(readTimeContainer);
  metadata.appendChild(topicContainer);
  header.appendChild(metadata);

  const summary = document.createElement("md-block");
  summary.textContent = post.summary;

  const contentElement = document.createElement("md-block");
  contentElement.textContent = post.content;

  blogPost.appendChild(header);
  blogPost.appendChild(summary);
  blogPost.appendChild(contentElement);

  postContainer.appendChild(blogPost);

  const backButton = document.createElement("button");
  backButton.classList.add("button");
  backButton.textContent = "Back to Posts";
  backButton.onclick = fetchAndDisplayPosts;
  postContainer.appendChild(backButton);
}
