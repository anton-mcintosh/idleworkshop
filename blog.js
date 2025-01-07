async function fetchAndDisplayPosts() {
  try {
    const response = await fetch("https://idleworkshop.com/api/get-posts");

    if (!response.ok) {
      throw new Error(`Error: ${response.statusText}`);
    }

    const data = await response.json();
    console.log(data); // debug

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
      const metadata = document.createElement("p");
      metadata.textContent = `Published: ${new Date(post.date).toLocaleDateString()} | Read Time: ${post.readTime} | Topic: ${post.topic}`;
      header.appendChild(title);
      header.appendChild(metadata);

      // the nutshell sub-heading
      const nutshell = document.createElement("p");
      nutshell.textContent = post.nutshell;

      // the 1-2 paragraph summary
      const summary = document.createElement("p");
      summary.textContent = post.summary;

      //footer with tags and button
      const footer = document.createElement("footer");
      const tagList = document.createElement("ul");
      for (const tag of post.tags) {
        const tagElement = document.createElement("li");
        tagElement.textContent = tag;
        tagList.appendChild(tagElement);
      }

      const button = document.createElement("button");
      button.classList.add("button");
      button.textContent = "Check it out";
      button.onclick = () => {
        window.location.href = `/blog/${post.slug}`;
      };
      footer.appendChild(tagList);
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
