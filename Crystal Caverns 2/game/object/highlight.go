components {
  id: "pfx"
  component: "/game/object/highlight.particlefx"
}
components {
  id: "highlight1"
  component: "/game/object/highlight.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"highlight-left\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ADD\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/game/core/game.tilesource\"\n"
  "}\n"
  ""
  position {
    x: -8.0
  }
}
embedded_components {
  id: "co"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_TRIGGER\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"prox\"\n"
  "mask: \"player\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 8.0\n"
  "  data: 8.0\n"
  "  data: 10.0\n"
  "}\n"
  ""
}
embedded_components {
  id: "sprite1"
  type: "sprite"
  data: "default_animation: \"highlight-right\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ADD\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/game/core/game.tilesource\"\n"
  "}\n"
  ""
  position {
    x: 8.0
  }
}
